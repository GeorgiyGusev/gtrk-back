package core

import (
	"context"
	"errors"
	"github.com/GeorgiyGusev/gtrk-back/pkg/http_server/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"log/slog"
	"net"
	"net/http"
	"time"
)

func NewHttpServer(cfg *Config, middlewares []echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()
	for _, m := range middlewares {
		e.Use(m)
	}
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.AllowedOrigins,
		AllowHeaders: cfg.AllowedHeaders,
		AllowMethods: cfg.AllowedMethods,
	}))
	return e
}

func RunHttpServer(lc fx.Lifecycle, e *echo.Echo, cfg *Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", cfg.Address())
			if err != nil {
				slog.Error("cannot start server", "err", err.Error(), constants.HttpServerTag)
				return err
			}
			e.Listener = listener
			slog.Info("starting server", constants.HttpServerTag, "address", cfg.Address())
			go func() {
				err := e.Start("")
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					slog.Error("cannot start server, force exit", "err", err.Error(), constants.HttpServerTag)
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("shutting down", constants.HttpServerTag)
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			return e.Shutdown(ctx)
		},
	})
}
