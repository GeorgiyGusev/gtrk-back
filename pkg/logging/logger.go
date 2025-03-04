package logging

import (
	"context"
	"errors"
	"log/slog"
	"os"
)

const (
	errLogKey = "err"
)

func ErrorField(err error) slog.Attr {
	return slog.String(errLogKey, err.Error())
}

type errorWithLogCtx struct {
	next error
	ctx  logCtx
}

func (e *errorWithLogCtx) Error() string {
	return e.next.Error()
}

func (e *errorWithLogCtx) Unwrap() error {
	return e.next
}

func WrapError(ctx context.Context, err error) error {
	c := logCtx{}
	if x, ok := ctx.Value(key).(logCtx); ok {
		c = x
	}
	return &errorWithLogCtx{
		next: err,
		ctx:  c,
	}
}

type HandlerMiddlware struct {
	next slog.Handler
}

func NewHandlerMiddleware(next slog.Handler) *HandlerMiddlware {
	return &HandlerMiddlware{next: next}
}

func (h *HandlerMiddlware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HandlerMiddlware{next: h.next.WithAttrs(attrs)} // не забыть обернуть, но осторожно
}

func (h *HandlerMiddlware) WithGroup(name string) slog.Handler {
	return &HandlerMiddlware{next: h.next.WithGroup(name)} // не забыть обернуть, но осторожно
}

func (h *HandlerMiddlware) Enabled(ctx context.Context, rec slog.Level) bool {
	return h.next.Enabled(ctx, rec)
}

type logCtx struct {
	UserID    *string
	UserEmail *string
}

func (h *HandlerMiddlware) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(key).(logCtx); ok {
		if c.UserID != nil {
			rec.Add("userID", c.UserID)
		}
		if c.UserEmail != nil {
			rec.Add("userEmail", c.UserEmail)
		}

	}
	return h.next.Handle(ctx, rec)
}

type keyType int

const key = keyType(0)

func WithLogUserID(ctx context.Context, userID string) context.Context {
	if c, ok := ctx.Value(key).(logCtx); ok {
		c.UserID = &userID
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, logCtx{UserID: &userID})
}

func WithLogUserEmail(ctx context.Context, email string) context.Context {
	if c, ok := ctx.Value(key).(logCtx); ok {
		c.UserEmail = &email
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, logCtx{UserEmail: &email})
}

// Пример с мутацией данных епта масло
//func WithLogPhone(ctx context.Context, phone string) context.Context {
//	if len(phone) > 4 {
//		phone = strings.Repeat("*", len(phone)-4) + phone[len(phone)-4:]
//	}
//	if c, ok := ctx.Value(key).(logCtx); ok {
//		c.Phone = phone
//		return context.WithValue(ctx, key, c)
//	}
//	return context.WithValue(ctx, key, logCtx{Phone: phone})
//}
// -----------------------------------------------

func ErrorCtx(ctx context.Context, err error) context.Context {
	var e *errorWithLogCtx
	if errors.As(err, &e) { // в реальной жизни используйте error.As
		return context.WithValue(ctx, key, e.ctx)
	}
	return ctx
}

// -----------------------------------------------

func InitLogging() {
	handler := slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	handler = NewHandlerMiddleware(handler)
	slog.SetDefault(slog.New(handler))
}
