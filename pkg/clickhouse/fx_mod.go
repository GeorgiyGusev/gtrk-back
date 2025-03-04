package clickhouse

import "go.uber.org/fx"

var Module = fx.Module(
	"clickhouse",
	fx.Provide(LoadConfig, NewClickhouseConn),
)
