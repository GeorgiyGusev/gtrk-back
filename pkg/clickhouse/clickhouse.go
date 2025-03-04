package clickhouse

import (
	"context"
	// Explicitly import the ClickHouse driver
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func NewClickhouseConn(lc fx.Lifecycle, cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("clickhouse", cfg.GetDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.DbMaxConns)

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return db.Close()
		},
	})
	return db, nil
}
