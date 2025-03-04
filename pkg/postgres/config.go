package postgres

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbName     string `env:"POSTGRES_NAME" env-default:"postgres"`
	DbUser     string `env:"POSTGRES_USER" env-default:"postgres"`
	DbPass     string `env:"POSTGRES_PASS" env-default:"postgres"`
	DbHost     string `env:"POSTGRES_HOST" env-default:"localhost"`
	DbPort     int    `env:"POSTGRES_PORT" env-default:"5432"`
	DbSslMode  string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
	DbMaxConns int    `env:"POSTGRES_MAX_CONNECTIONS" env-default:"0"`
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName, c.DbSslMode)
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
