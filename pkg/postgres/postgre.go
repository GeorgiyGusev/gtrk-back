package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

type PostressConn struct {
	DB *sqlx.DB
}

func NewPostgresConn(lc fx.Lifecycle, cfg *Config) (*PostressConn, error) {
	db, err := sqlx.Connect("postgres", cfg.GetDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.DbMaxConns)

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return db.Close()
		},
	})
	return &PostressConn{
		DB: db,
	}, nil
}
