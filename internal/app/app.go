package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"trade-union-service/internal/config"
	"trade-union-service/internal/domains/telegram"
	"trade-union-service/internal/domains/users"

	setupEchotron "trade-union-service/internal/setup/echotron"
	setupFiber "trade-union-service/internal/setup/fiber"
	setupLogger "trade-union-service/internal/setup/logger"
	setupMongoDB "trade-union-service/internal/setup/mongodb"
	setupVictoriaMetrics "trade-union-service/internal/setup/victoria-metrics"

	mongoDBMigrate "trade-union-service/internal/setup/mongodb/migrate"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type setup struct {
	fiber *setupFiber.Fiber
	// redis    *setupRedis.Redis
	mongodb  *setupMongoDB.MongoDB
	echotron *setupEchotron.Echotron
	metrics  *setupVictoriaMetrics.VictoriaMetrics
}

type migrate struct {
	mongodb *mongoDBMigrate.Migrate
}

type app struct {
	log      *logrus.Logger
	cfg      *config.Config
	errGroup *errgroup.Group
	setup    setup
	migrate  migrate
	stopCh   chan os.Signal
	ctx      context.Context
	cancelFn context.CancelFunc
}

func NewApp() *app {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background()) // Base app context

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	logger := setupLogger.NewLogger(&cfg.Logger)

	errGroup, ctx := errgroup.WithContext(ctx)

	// redis := setupRedis.NewRedis(&cfg.Redis, logger)

	mongodb, err := setupMongoDB.NewMongoDB(ctx, &cfg.MongoDB, logger)
	if err != nil {
		panic(err)
	}

	mongodbMigrate, err := mongoDBMigrate.NewMigrate(mongodb, logger)
	if err != nil {
		panic(err)
	}

	fiber := setupFiber.NewFiber(&cfg.Fiber, logger)

	echotron, err := setupEchotron.NewEchotron(&cfg.Echotron, logger)
	if err != nil {
		panic(err)
	}

	metrics := setupVictoriaMetrics.NewVictoriaMetrics(&cfg.VictoriaMetrics, logger)

	return &app{
		log:      logger,
		cfg:      cfg,
		errGroup: errGroup,
		stopCh:   stopCh,
		ctx:      ctx,
		cancelFn: cancel,
		setup: setup{
			fiber:   fiber,
			mongodb: mongodb,
			// redis:    redis,
			echotron: echotron,
			metrics:  metrics,
		},
		migrate: migrate{
			mongodb: mongodbMigrate,
		},
	}
}

func (a *app) Run() error {
	// Run migrations
	{
		if err := a.migrate.mongodb.Run(); err != nil {
			return err
		}
	}

	// Initialize domains
	{
		telegram.NewDomain(a.setup.mongodb, a.setup.echotron, a.setup.metrics, a.log)
		users.NewDomain(a.setup.mongodb, a.setup.fiber, a.setup.metrics, a.log)
	}

	// Handle stop program
	a.errGroup.Go(func() error {
		sig := <-a.stopCh
		a.log.Infof("Got %s signal. Aborting...\n", sig)
		a.cancelFn()
		close(a.stopCh)
		return nil
	})

	// Run fiber http server
	a.errGroup.Go(func() error {
		if err := a.setup.fiber.Run(a.ctx); err != nil {
			return err
		}
		return nil
	})

	// Run metrics
	a.errGroup.Go(func() error {
		if err := a.setup.metrics.Run(); err != nil {
			return err
		}
		return nil
	})

	// Wait error from group of goroutines
	if err := a.errGroup.Wait(); err != nil {
		a.log.Error("app: Run - g.Wait error: ", err.Error())
		return err
	}

	return nil
}
