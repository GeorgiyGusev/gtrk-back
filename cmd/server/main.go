package main

import (
	"github.com/GeorgiyGusev/gtrk-back/internal/stats"
	"github.com/GeorgiyGusev/gtrk-back/internal/user_subs"
	"github.com/GeorgiyGusev/gtrk-back/pkg/clickhouse"
	"github.com/GeorgiyGusev/gtrk-back/pkg/core_sp"
	"github.com/GeorgiyGusev/gtrk-back/pkg/grpc_server"
	"github.com/GeorgiyGusev/gtrk-back/pkg/logging"
	"github.com/GeorgiyGusev/gtrk-back/pkg/postgres"
	"github.com/GeorgiyGusev/gtrk-back/pkg/validation"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func main() {
	logging.InitLogging()
	app := fx.New(
		// setting logger
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.SlogLogger{
				Logger: slog.Default(),
			}
		}),

		// including platform libs here
		validation.Module,
		//http_server.Module,
		// Uncomment if need to use grpc server
		grpc_server.Module,
		postgres.Module,
		clickhouse.Module,
		db_sp_call.Module,

		// domains
		user_subs.Module,
		stats.Module,
	)
	app.Run()
}
