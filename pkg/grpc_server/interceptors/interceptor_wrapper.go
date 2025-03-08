package interceptors

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

const (
	UnaryServerInterceptorGroup  = `group:"unaryServerInterceptors"`
	StreamServerInterceptorGroup = `group:"streamServerInterceptors"`
)

type UnaryInterceptor interface {
	UnaryWrap() grpc.UnaryServerInterceptor
}

type StreamInterceptor interface {
	StreamWrap() grpc.StreamServerInterceptor
}

func BuildUnaryInterceptorChaim(interceptors []UnaryInterceptor) []grpc.UnaryServerInterceptor {
	var chain []grpc.UnaryServerInterceptor
	for _, interceptor := range interceptors {
		chain = append(chain, interceptor.UnaryWrap())
	}
	return chain
}

func BuildStreamInterceptorChaim(interceptors []StreamInterceptor) []grpc.StreamServerInterceptor {
	var chain []grpc.StreamServerInterceptor
	for _, interceptor := range interceptors {
		chain = append(chain, interceptor.StreamWrap())
	}
	return chain
}

func AsUnaryInterceptor(interceptor UnaryInterceptor) any {
	return fx.Annotate(
		interceptor,
		fx.As(new(UnaryInterceptor)),
		fx.ResultTags(UnaryServerInterceptorGroup),
	)
}

func AsStreamInterceptor(interceptor StreamInterceptor) any {
	return fx.Annotate(
		interceptor,
		fx.As(new(StreamInterceptor)),
		fx.ResultTags(StreamServerInterceptorGroup),
	)
}
