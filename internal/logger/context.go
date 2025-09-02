package logger

type ContextKey string

const (
	ContextKeyRequestID ContextKey = "request_id"
	ContextKeyError     ContextKey = "error"
)
