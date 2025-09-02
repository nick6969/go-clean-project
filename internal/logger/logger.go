package logger

import "context"

// Logger defines the interface for logging with context awareness.
type Logger interface {
	With(ctx context.Context) Logger
	WithAdditionalFields(fields map[string]any) Logger

	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
}
