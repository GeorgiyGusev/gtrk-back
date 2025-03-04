package user_subs

import "go.uber.org/fx"

var Module = fx.Module(
	"user_subs",
	fx.Provide(NewHandler),
	fx.Provide(fx.Annotate(NewRepoImpl, fx.As(new(Repo)))),
	fx.Invoke(RegisterHandlers),
)
