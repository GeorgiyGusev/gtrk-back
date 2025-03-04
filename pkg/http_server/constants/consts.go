package constants

import "log/slog"

const (
	HttpMiddlewareGroup = `group:"httpMiddleware"`
)

var (
	HttpServerTag = slog.String("server", "http_server")
)
