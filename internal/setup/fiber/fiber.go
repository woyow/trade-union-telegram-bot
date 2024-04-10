package fiber

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v3"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	addressSeparator = ":"
)

type Fiber struct {
	cfg *Config
	App *fiber.App
	log *logrus.Logger
}

// NewFiber - .
func NewFiber(cfg *Config, log *logrus.Logger) *Fiber {
	corsMiddleware := getCorsMiddleware(&cfg.Handler.CORS)
	app := getFiberApp(cfg)
	app.Use(corsMiddleware)

	return &Fiber{
		App: app,
		cfg: cfg,
		log: log,
	}
}

func (f *Fiber) Run(ctx context.Context) error {
	address := f.cfg.Host + addressSeparator + f.cfg.Port

	if err := f.App.Listen(address, fiber.ListenConfig{
		ListenerNetwork:       "tcp4",
		GracefulContext:       ctx,
		DisableStartupMessage: false,
		EnablePrefork:         false,
		EnablePrintRoutes:     true,
	}); err != nil {
		f.log.Error("fiber: Run - f.app.Listen error: ", err.Error())
		return err
	}
	return nil
}

func getCorsMiddleware(cfg *CORS) fiber.Handler {
	return cors.New(cors.Config{
		Next:                nil,
		AllowOriginsFunc:    nil,
		AllowOrigins:        strings.Join(cfg.AllowOrigins, ","),
		AllowMethods:        strings.Join(cfg.AllowMethods, ","),
		AllowHeaders:        strings.Join(cfg.AllowHeaders, ","),
		AllowCredentials:    cfg.AllowCredentials,
		ExposeHeaders:       "",
		MaxAge:              cfg.MaxAge,
		AllowPrivateNetwork: false,
	})
}

func getFiberApp(cfg *Config) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:         cfg.AppName,
		ReadTimeout:     time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:     time.Duration(cfg.IdleTimeout) * time.Second,
		ReadBufferSize:  cfg.ReadBufferSize,
		WriteBufferSize: cfg.WriteBufferSize,
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
	})
}