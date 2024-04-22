package http

import (
	"trade-union-service/internal/setup/http/client"
)

type Config struct {
	Client client.Config `yaml:"client"`
}
