package fiber

type Config struct {
	Handler         Handler `yaml:"handler"`
	AppName         string  `yaml:"app_name"`
	Host            string  `yaml:"host"`
	Port            string  `yaml:"port"`
	ReadTimeout     int64   `yaml:"read_timeout"`
	WriteTimeout    int64   `yaml:"write_timeout"`
	IdleTimeout     int64   `yaml:"idle_timeout"`
	ReadBufferSize  int     `yaml:"read_buffer_size"`
	WriteBufferSize int     `yaml:"write_buffer_size"`
}

type Handler struct {
	CORS CORS `yaml:"cors"`
}

// CORS - CORS parameters
type CORS struct {
	AllowMethods     []string `yaml:"allow_methods"`
	AllowHeaders     []string `yaml:"allow_headers"`
	AllowOrigins     []string `yaml:"allow_origins"`
	MaxAge           int      `yaml:"max_age"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	AllowAllOrigins  bool     `yaml:"allow_all_origins"`
}