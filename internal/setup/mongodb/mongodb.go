package mongodb

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	setupLoggingKey   = "setup"
	setupLoggingValue = "mongodb"

	proto            = "mongodb://"
	addressSeparator = ":"
	authMechanism    = "SCRAM-SHA-1"
)

type MongoDB struct {
	Client *mongo.Client
	Config *Config
	Cancel context.CancelFunc
}

func NewMongoDB(ctx context.Context, cfg *Config, log *logrus.Logger) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	mongoURI := proto + cfg.Host + addressSeparator + cfg.Port

	log.WithField(setupLoggingKey, setupLoggingValue).
		Debug("mongoURI: ", mongoURI)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI).SetAuth(options.Credential{
		AuthMechanism: authMechanism,
		AuthSource:    cfg.AuthSource,
		Username:      cfg.Username,
		Password:      cfg.Password,
	}))
	if err != nil {
		log.WithField(setupLoggingKey, setupLoggingValue).
			Error("NewMongoDB - mongo.Connect error: ", err.Error())

		cancel()

		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.WithField(setupLoggingKey, setupLoggingValue).
			Error("NewMongoDB - client.Ping error: ", err.Error())

		cancel()

		return nil, err
	}

	return &MongoDB{
		Client: client,
		Config: cfg,
		Cancel: cancel,
	}, nil
}
