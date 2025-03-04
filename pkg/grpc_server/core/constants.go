package core

import "log/slog"

const (
	UnaryServerInterceptorGroup  = `group:"unaryServerInterceptors"`
	StreamServerInterceptorGroup = `group:"streamServerInterceptors"`
)

var (
	grpcServerTag = slog.String("server", "grpc_server")
)
