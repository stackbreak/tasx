package web

import "github.com/spf13/viper"

type Config struct {
	*viper.Viper
}

func NewConfig() *Config {
	return &Config{viper.New()}
}

func (c *Config) Init() error {
	c.AddConfigPath("configs")
	c.SetConfigName("config")
	return c.ReadInConfig()
}
