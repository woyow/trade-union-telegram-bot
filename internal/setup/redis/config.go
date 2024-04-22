package redis

type Config struct {
	Password string `env:"REDIS_PASSWORD"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
}
