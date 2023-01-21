package config

const api = true

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) ApiActive() bool {
	return api
}
