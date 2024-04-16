package victoria_metrics

type Config struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	PushInterval   int64  `yaml:"push_interval"`
	MetricsEnabled bool   `yaml:"metrics_enabled"`
}
