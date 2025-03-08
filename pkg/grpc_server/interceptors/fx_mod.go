package interceptors

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Module(
	"grpc_interceptors",
	fx.Provide(
		fx.Annotate(NewUnaryLoggingInterceptor, fx.As(new(grpc.UnaryServerInterceptor)), fx.ResultTags(UnaryServerInterceptorGroup)),
		fx.Annotate(NewStreamLoggingInterceptor, fx.As(new(grpc.StreamServerInterceptor)), fx.ResultTags(StreamServerInterceptorGroup)),
		fx.Annotate(NewValidationInterceptor, fx.As(new(grpc.UnaryServerInterceptor)), fx.ResultTags(UnaryServerInterceptorGroup)),
		fx.Annotate(func(ui []UnaryInterceptor) []UnaryInterceptor { return ui }, fx.ParamTags(UnaryServerInterceptorGroup)),
		fx.Annotate(func(si []StreamInterceptor) []StreamInterceptor { return si }, fx.ParamTags(StreamServerInterceptorGroup)),
	),
)
