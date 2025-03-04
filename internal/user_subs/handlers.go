package user_subs

import (
	"context"
	"fmt"
	"github.com/GeorgiyGusev/gtrk-back/gen/proto/users_subs/v1"
	"github.com/GeorgiyGusev/gtrk-back/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Repo interface {
	Create(ctx context.Context, request *users_subs_gen_v1.CreateUserSubRequest) error
	GetAll(ctx context.Context) ([]*users_subs_gen_v1.CreateUserSubResponse, error)
}

type Handler struct {
	repo Repo
	users_subs_gen_v1.UnimplementedUsersSubsServiceServer
}

func RegisterHandlers(srv *grpc.Server, handlers *Handler) {
	users_subs_gen_v1.RegisterUsersSubsServiceServer(srv, handlers)
}

func NewHandler(repo Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateUserSub(ctx context.Context, request *users_subs_gen_v1.CreateUserSubRequest) (*users_subs_gen_v1.CreateUserSubResponse, error) {

	if err := h.repo.Create(ctx, request); err != nil {
		slog.ErrorContext(logging.ErrorCtx(ctx, err), "create user sub failed", "err", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("create user sub failed: %v", err))
	}
	return &users_subs_gen_v1.CreateUserSubResponse{}, nil
}
