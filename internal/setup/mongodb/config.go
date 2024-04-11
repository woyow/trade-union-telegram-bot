package mongodb

type Config struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Username   string `env:"MONGO_USERNAME"`
	Password   string `env:"MONGO_PASSWORD"`
	AuthSource string `env:"MONGO_DATABASE"`
}
