package config

type Config struct {
	Port int
}

func LoadConfig() (*Config, error) {
	return &Config{
		Port: 8080,
	}, nil
}
