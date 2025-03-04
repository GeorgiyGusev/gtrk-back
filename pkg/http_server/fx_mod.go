package http_server

import (
	"github.com/GeorgiyGusev/gtrk-back/pkg/http_server/constants"
	"github.com/GeorgiyGusev/gtrk-back/pkg/http_server/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http_server",
	fx.Provide(
		core.LoadConfig,
		fx.Annotate(core.NewHttpServer, fx.ParamTags(``, constants.HttpMiddlewareGroup)),
	),
	fx.Invoke(core.RunHttpServer),
)
