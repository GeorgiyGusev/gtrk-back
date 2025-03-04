package interceptors

import (
	"github.com/GeorgiyGusev/gtrk-back/pkg/grpc_server/core"
	"go.uber.org/fx"
)

func AsUnaryServerInterceptor(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(core.UnaryServerInterceptorGroup),
	)
}
func AsStreamServerInterceptor(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(core.StreamServerInterceptorGroup),
	)
}
