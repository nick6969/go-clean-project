package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/nick6969/go-clean-project/internal/config"
)

type Slogger struct {
	logger *slog.Logger
	ctx    context.Context
}

func NewSLogger(ctx context.Context, cfg config.LoggerConfig) *Slogger {
	return &Slogger{
		logger: slog.New(getHandler(cfg)),
		ctx:    ctx,
	}
}

func (l *Slogger) GetDatabaseLogger() *slog.Logger {
	return l.logger.With("component", "database")
}

func (l *Slogger) With(ctx context.Context) Logger {
	attrs := extractSlogAttributes(ctx)
	return &Slogger{
		logger: l.logger.With(attrs...),
		ctx:    ctx,
	}
}

func (l *Slogger) WithAdditionalFields(fields map[string]any) Logger {
	return &Slogger{
		logger: l.logger.With(l.getAttrs(fields, nil)...),
		ctx:    l.ctx,
	}
}

func (l *Slogger) Debug(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *Slogger) Info(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *Slogger) Warn(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *Slogger) Error(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func (l *Slogger) getAttrs(fields map[string]any, err error) []any {
	var attrs []any
	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}
	if err != nil {
		attrs = append(attrs, slog.Any(string(ContextKeyError), err))
	}
	return attrs
}

// help functions
func getLevelFromString(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func getHandler(cfg config.LoggerConfig) slog.Handler {
	level := getLevelFromString(cfg.Level)
	if strings.ToLower(cfg.Format) == "json" {
		return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
}

func extractSlogAttributes(ctx context.Context) []any {
	attrs := make([]any, 0)

	if requestID := ctx.Value(ContextKeyRequestID); requestID != nil {
		if requestIDStr, ok := requestID.(string); ok && requestIDStr != "" {
			attrs = append(attrs, slog.Any(string(ContextKeyRequestID), requestIDStr))
		}
	}

	return attrs
}
