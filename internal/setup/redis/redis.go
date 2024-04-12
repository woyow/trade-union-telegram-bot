package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// Redis - Redis storage.
type Redis struct {
	client *redis.Client
	log    *logrus.Logger
}

// NewRedis - Returns *Redis.
func NewRedis(cfg *Config, log *logrus.Logger) *Redis {
	client := getRedisClient(cfg)

	return &Redis{
		client: client,
		log:    log,
	}
}

const (
	addressSeparator = ":"
)

// getRedisClient - Returns *redis.Client.
func getRedisClient(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Host + addressSeparator + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
