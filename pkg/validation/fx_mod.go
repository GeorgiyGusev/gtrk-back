package validation

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"validators",
	fx.Provide(
		protovalidate.New,
		validator.New,
	),
)
