package web

import (
	"github.com/caarlos0/env/v6"
)

type EnvVars struct {
	AppPort string `env:"APP_PORT,required"`

	DbUser    string `env:"DB_USER,required"`
	DbPass    string `env:"DB_PASS,required"`
	DbName    string `env:"DB_NAME,required"`
	DbHost    string `env:"DB_HOST,required"`
	DbPort    string `env:"DB_PORT,required"`
	DbSslMode string `env:"DB_SSLMODE,required"`

	TokenSecret string `env:"TOKEN_SECRET,required"`
}

type Config struct {
	Env *EnvVars
}

func NewConfig() *Config {
	return &Config{&EnvVars{}}
}

func (c *Config) LoadEnv() error {
	if err := env.Parse(c.Env); err != nil {
		return err
	}

	return nil
}
