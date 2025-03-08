package core

import (
	"context"
	"github.com/GeorgiyGusev/gtrk-back/pkg/grpc_server/interceptors"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type UnaryServerInterceptors []grpc.UnaryServerInterceptor
type StreamServerInterceptors []grpc.StreamServerInterceptor

func NewGrpcServer(unaryInterceptors []interceptors.UnaryInterceptor, streamInterceptors []interceptors.StreamInterceptor) *grpc.Server {
	slog.Info("Initializing server", grpcServerTag)
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.BuildUnaryInterceptorChaim(unaryInterceptors)...),
		grpc.ChainStreamInterceptor(interceptors.BuildStreamInterceptorChaim(streamInterceptors)...),
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(server)
	return server
}

func RunGrpcServer(lc fx.Lifecycle, srv *grpc.Server, cfg *Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			slog.Info("starting server", grpcServerTag, "address", cfg.Address())
			listener, err := net.Listen("tcp", cfg.Address())
			if err != nil {
				slog.Error("cannot start server", "error", err.Error(), grpcServerTag)
				return err
			}
			go func() {
				err := srv.Serve(listener)
				if err != nil {
					slog.Error("cannot start server", "error", err.Error(), grpcServerTag)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("shutting down", grpcServerTag)
			srv.GracefulStop()
			return nil
		},
	})
}
