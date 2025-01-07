package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

// Logger is a simple logger instance
var Logger = slog.New(slog.NewTextHandler(os.Stdout))

// Middleware is an empty middleware for the handler
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic goes here

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
type CustomHandler struct {
	logger slog.Handler
}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handler logic goes here
}

func (h *CustomHandler) Log(record slog.Record) error {
	// Log method logic goes here
	return nil
}

func (h *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.logger.Enabled(ctx, level)
}

func (h *CustomHandler) Handle(ctx context.Context, record slog.Record) error {
	return h.logger.Handle(ctx, record)
}

func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CustomHandler{
		logger: h.logger.WithAttrs(attrs),
	}
}
