package logger

// Config - Logger config
type Config struct {
	Level            string `yaml:"log_level"`
	DisableTimestamp bool   `yaml:"disable_timestamp"`
	FullTimestamp    bool   `yaml:"full_timestamp"`
}
