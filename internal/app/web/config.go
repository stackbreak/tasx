package web

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type EnvVars struct {
	DbUser    string `env:"DB_USER,required"`
	DbPass    string `env:"DB_PASS,required"`
	DbName    string `env:"DB_NAME,required"`
	DbHost    string `env:"DB_HOST,required"`
	DbPort    string `env:"DB_PORT,required"`
	DbSslMode string `env:"DB_SSLMODE,required"`

	TokenSecret string `env:"TOKEN_SECRET,required"`
}

type Config struct {
	File *viper.Viper
	Env  *EnvVars
}

func NewConfig() *Config {
	return &Config{viper.New(), &EnvVars{}}
}

func (c *Config) LoadFile() error {
	c.File.AddConfigPath("configs")
	c.File.SetConfigName("web")
	return c.File.ReadInConfig()
}

func (c *Config) LoadEnv() error {
	if err := godotenv.Load("./configs/local.env"); err != nil {
		return err
	}

	if err := env.Parse(c.Env); err != nil {
		return err
	}

	return nil
}
