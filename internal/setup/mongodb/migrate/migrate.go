package migrate

import (
	"errors"

	setupMongo "trade-union-service/internal/setup/mongodb"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

type Migrate struct {
	m   *migrate.Migrate
	log *logrus.Logger
}

const (
	setupLoggingKey   = "setup"
	setupLoggingValue = "mongodb migrate"

	sourceURL    = "file://db/mongodb/migrations"
	databaseName = "mongodb"
)

func NewMigrate(setupMongo *setupMongo.MongoDB, log *logrus.Logger) (*Migrate, error) {
	driver, err := mongodb.WithInstance(setupMongo.Client, &mongodb.Config{
		DatabaseName:    setupMongo.Config.AuthSource,
		TransactionMode: false,
	})
	if err != nil {
		log.WithField(setupLoggingKey, setupLoggingValue).
			Error("NewMigrate - mongodb.WithInstance error: ", err.Error())
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, databaseName, driver)
	if err != nil {
		log.WithField(setupLoggingKey, setupLoggingValue).
			Error("NewMigrate - migrate.NewWithDatabaseInstance error: ", err.Error())
		return nil, err
	}

	return &Migrate{
		m:   m,
		log: log,
	}, nil
}

func (m *Migrate) Run() error {
	if err := m.m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			m.log.WithField(setupLoggingKey, setupLoggingValue).
				Info("Run - no change")

			return nil
		}

		m.log.WithField(setupLoggingKey, setupLoggingValue).
			Fatal("Run - m.m.Up error: ", err.Error())

		return err
	}

	return nil
}
