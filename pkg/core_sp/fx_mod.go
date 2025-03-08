package db_sp_call

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"db_sp_call",
	fx.Provide(fx.Annotate(NewDBCall)),
)
