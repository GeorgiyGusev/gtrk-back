package main

import (
	"context"
	"log/slog"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewServer,
			fx.Annotate(NewCORS, fx.As(new(Middleware)), fx.ResultTags(`group:"mw"`)),
			fx.Annotate(NewLogger, fx.As(new(Middleware)), fx.ResultTags(`group:"mw"`)),
			fx.Annotate(func(mm []Middleware) []Middleware { return mm }, fx.ParamTags(`group:"mw"`)),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

type Middleware interface {
	UnWrap(http.Handler) http.Handler
}

type CORS struct {
}

func NewCORS() *CORS {
	return &CORS{}
}

func (c *CORS) UnWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("cors")
		h.ServeHTTP(w, r)
	})
}

type Logger struct{}

func NewLogger( /* deps1, deps2, deps3, deps4 */ ) *Logger {
	return new(Logger)
}

func (l *Logger) UnWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("logger")
		h.ServeHTTP(w, r)
	})
}

func NewServer(middlewares []Middleware, lc fx.Lifecycle) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	var h http.Handler = mux
	for _, m := range middlewares {
		h = m.UnWrap(h)
	}
	s := &http.Server{
		Addr:    ":http",
		Handler: h,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go s.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})
	return s
}
