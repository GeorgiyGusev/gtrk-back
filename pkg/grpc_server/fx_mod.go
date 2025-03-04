package grpc_server

import (
	"github.com/GeorgiyGusev/gtrk-back/pkg/grpc_server/core"
	"github.com/GeorgiyGusev/gtrk-back/pkg/grpc_server/interceptors"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"grpc_server",
	fx.Provide(
		core.LoadConfig,
		fx.Annotate(
			core.NewGrpcServer,
			fx.ParamTags(core.UnaryServerInterceptorGroup, core.StreamServerInterceptorGroup),
		),
		interceptors.AsUnaryServerInterceptor(interceptors.NewValidateInterceptor),
		interceptors.AsUnaryServerInterceptor(interceptors.NewUnaryLoggingInterceptor),
		interceptors.AsStreamServerInterceptor(interceptors.NewStreamLoggingInterceptor),
	),
	fx.Invoke(
		core.RunGrpcServer,
	),
)
