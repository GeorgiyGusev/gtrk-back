package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"net"
	"strconv"
)

type Config struct {
	Host string `env-default:"0.0.0.0" env:"GRPC_SERVER_HOST"`
	Port int    `env-default:"50051" env:"GRPC_SERVER_PORT"`
}

func (cfg *Config) Address() string {
	return net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
}

func LoadConfig(v *validator.Validate) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	if err := v.Struct(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
