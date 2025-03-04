package interceptors

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"log/slog"
)

func InterceptorLoggerAdapter() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		slog.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func NewUnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall, logging.PayloadSent, logging.PayloadReceived),
	}
	return logging.UnaryServerInterceptor(InterceptorLoggerAdapter(), opts...)
}

func NewStreamLoggingInterceptor() grpc.StreamServerInterceptor {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall, logging.PayloadSent, logging.PayloadReceived),
	}
	return logging.StreamServerInterceptor(InterceptorLoggerAdapter(), opts...)
}
