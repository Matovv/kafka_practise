package configs

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BindAddr         string   `env:"BIND_ADDR"    toml:"bind_addr"    env-default:":8000"`
	DatabaseURI      string   `env:"DATABASE_URI" toml:"database_uri"`
}

func ParseConfig(configPath string) (*Config, error) {
	config := &Config{}

	var err error

	if configPath != "" {
		err = cleanenv.ReadConfig(configPath, config)
	} else {
		err = cleanenv.ReadEnv(config)
	}

	if err != nil {
		return nil, errors.New("failed to parse config")
	}
	return config, nil
}