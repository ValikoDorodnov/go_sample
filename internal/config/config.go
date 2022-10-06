package config

import (
	"fmt"

	"github.com/ValikoDorodnov/go_sample/pkg/config"
	"github.com/joho/godotenv"
)

type Config struct {
	Rest RestConfig `mapstructure:",squash"`
	Db   DbConfig   `mapstructure:",squash"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func InitConfig() *Config {
	conf := &Config{}
	err := config.Unmarshal(conf)
	if err != nil {
		panic(fmt.Sprintf("Can't create config: %v", err))
	}

	return conf
}
