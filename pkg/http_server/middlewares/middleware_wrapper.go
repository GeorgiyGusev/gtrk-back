package middlewares

import (
	"github.com/GeorgiyGusev/gtrk-back/pkg/http_server/constants"
	"go.uber.org/fx"
)

func AsHttpInterceptor(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(constants.HttpMiddlewareGroup),
	)
}
