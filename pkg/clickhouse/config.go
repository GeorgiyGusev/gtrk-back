package clickhouse

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbName     string `env:"CLICKHOUSE_NAME" env-default:"default"`
	DbUser     string `env:"CLICKHOUSE_USER" env-default:"clickhouse"`
	DbPass     string `env:"CLICKHOUSE_PASS" env-default:"clickhouse"`
	DbHost     string `env:"CLICKHOUSE_HOST" env-default:"localhost"`
	DbPort     int    `env:"CLICKHOUSE_PORT" env-default:"9000"`
	DbMaxConns int    `env:"CLICKHOUSE_MAX_CONNECTIONS" env-default:"0"`
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("clickhouse://%s:%s@%s:%d/%s", c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName)
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
