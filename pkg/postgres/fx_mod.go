package postgres

import "go.uber.org/fx"

const (
	PostgreResultTag = `group:"postgres"`
)

var Module = fx.Module(
	"postgres",
	fx.Provide(LoadConfig, NewPostgresConn),
)
