package migrate

import (
	"errors"

	"github.com/sirupsen/logrus"
	setupMongo "trade-union-service/internal/setup/mongodb"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrate struct {
	m   *migrate.Migrate
	log *logrus.Logger
}

const (
	sourceURL    = "file://db/mongodb/migrations"
	databaseName = "mongodb"
)

func NewMigrate(setupMongo *setupMongo.MongoDB, log *logrus.Logger) (*Migrate, error) {
	driver, err := mongodb.WithInstance(setupMongo.Client, &mongodb.Config{
		DatabaseName:    setupMongo.Config.AuthSource,
		TransactionMode: false,
	})
	if err != nil {
		log.Error("mongodb migrate: NewMigrate - mongodb.WithInstance error: ", err.Error())
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, databaseName, driver)
	if err != nil {
		log.Error("mongodb migrate: NewMigrate - migrate.NewWithDatabaseInstance error: ", err.Error())
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
			m.log.Debug("mongodb migrate: Run - no change")
			return nil
		}

		m.log.Fatal("mongodb migrate: Run - m.m.Up error: ", err.Error())
		return err
	}
	return nil
}
