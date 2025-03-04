package user_subs

import (
	"context"
	users_subs_gen_v1 "github.com/GeorgiyGusev/gtrk-back/gen/proto/users_subs/v1"
	db_sp_call "github.com/GeorgiyGusev/gtrk-back/pkg/core_sp"
	"github.com/GeorgiyGusev/gtrk-back/pkg/logging"
	"log/slog"
)

type RepoImpl struct {
	call *db_sp_call.DBCall
}

type Resp struct {
}

func NewRepoImpl(call *db_sp_call.DBCall) *RepoImpl {
	return &RepoImpl{call: call}
}

func (r *RepoImpl) Create(ctx context.Context, request *users_subs_gen_v1.CreateUserSubRequest) error {

	var res interface{}

	dbErr, err := r.call.CallFunction(&res, "user_subs", "create_user_sub", request)
	if err != nil {
		slog.ErrorContext(logging.ErrorCtx(ctx, err), "user_stubs call procedure failed", logging.ErrorField(err))
		return err
	}
	if dbErr.HasErr() {
		slog.ErrorContext(logging.ErrorCtx(ctx, err), "user_stubs call procedure failed", logging.ErrorField(dbErr))
		return dbErr
	}
	return nil
}

func (r *RepoImpl) GetAll(ctx context.Context) ([]*users_subs_gen_v1.CreateUserSubResponse, error) {
	//TODO implement me
	panic("implement me")
}
