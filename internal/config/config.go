package config

import (
	"log"
	"os"
	"path/filepath"

	"trade-union-service/internal/setup/app"
	"trade-union-service/internal/setup/echotron"
	"trade-union-service/internal/setup/fiber"
	"trade-union-service/internal/setup/http"
	"trade-union-service/internal/setup/logger"
	"trade-union-service/internal/setup/mongodb"
	"trade-union-service/internal/setup/redis"
	victoriaMetrics "trade-union-service/internal/setup/victoria-metrics"

	goEnv "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const (
	configDir = "configs"
)

// Config - Aggregate configurations for application.
type Config struct {
	App             app.Config             `yaml:"app"`
	HTTP            http.Config            `yaml:"http"`
	Fiber           fiber.Config           `yaml:"fiber"`
	Logger          logger.Config          `yaml:"logger"`
	Redis           redis.Config           `yaml:"redis"`
	MongoDB         mongodb.Config         `yaml:"mongodb"`
	Echotron        echotron.Config        `yaml:"-"`
	VictoriaMetrics victoriaMetrics.Config `yaml:"victoria_metrics"`
}

// NewConfig - Returns *Config.
func NewConfig() (*Config, error) {
	var cfg Config

	if err := cfg.readEnv(); err != nil {
		return nil, err
	}

	if err := cfg.readConfigFile(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// readEnv - Read environment file (.env) and unmarshal data to config
func (cfg *Config) readEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("config: readEnv - godotenv.Load error: ", err.Error())
		return err
	}

	if _, err := goEnv.UnmarshalFromEnviron(cfg); err != nil {
		log.Fatal("config: readEnv - goEnv.UnmarshalFromEnviron error: ", err.Error())
		return err
	}

	return nil
}

// readConfigFile - Read yaml config file and unmarshal data to config
func (cfg *Config) readConfigFile() error {
	fileName := configDir + "/" + cfg.App.Env + ".yaml"

	filePath, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal("config: readConfigFile - filepath.Abs error: ", err.Error())
		return err
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("config: readConfigFile - os.ReadFile error: ", err.Error())
		return err
	}

	if err = yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatal("config: readConfigFile - yaml.Unmarshal error: ", err.Error())
		return err
	}

	return err
}
