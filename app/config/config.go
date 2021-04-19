package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr string `required:"true" split_words:"true"`
}

func New() (Config, error) {
	var cfg Config
	err := envconfig.Process("iban", &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, err
}
