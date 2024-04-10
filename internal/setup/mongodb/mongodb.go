package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	proto = "mongodb://"
	addressSeparator = ":"
)

type MongoDB struct {
	Client *mongo.Client
	Cancel context.CancelFunc
}

func NewMongoDB(ctx context.Context, cfg *Config) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(proto + cfg.Host + addressSeparator + cfg.Port))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoDB{
		Client: client,
		Cancel: cancel,
	}, nil
}
