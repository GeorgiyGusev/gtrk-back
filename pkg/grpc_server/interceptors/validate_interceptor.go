package interceptors

import (
	"context"
	"github.com/GeorgiyGusev/gtrk-back/pkg/logging"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log/slog"
)

func NewValidateInterceptor(validator *protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if msg, ok := req.(proto.Message); ok {
			if err := validator.Validate(msg); err != nil {
				slog.ErrorContext(ctx, "proto validation error", logging.ErrorField(err), "method", info.FullMethod)
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		// Продолжаем обработку
		return handler(ctx, req)
	}
}
