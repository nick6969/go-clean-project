package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/logger"
)

type RequestLogger struct {
	logger logger.Logger
}

func NewRequestLogger(logger logger.Logger) *RequestLogger {
	return &RequestLogger{
		logger: logger,
	}
}

func (l *RequestLogger) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := l.logger.With(c)

		start := time.Now()

		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		ip := getClientIP(c)

		l.Info(c, "Incoming request", map[string]any{
			"method":     method,
			"path":       path,
			"query":      query,
			"client_ip":  ip,
			"user_agent": c.Request.UserAgent(),
		})

		c.Next()

		latency := time.Since(start)

		l.Info(c, "Outgoing response", map[string]any{
			"method":    method,
			"path":      path,
			"query":     query,
			"client_ip": ip,
			"latency":   latency.String(),
			"status":    c.Writer.Status(),
			"body_size": c.Writer.Size(),
		})
	}
}

func getClientIP(c *gin.Context) string {
	// 這裏可以擴展以支援更多的 Forwarded Headers，如 X-Forwarded-For
	// 這是要根據實際部署環境來決定的
	// 這裏為了簡單起見，我們只使用 ClientIP 方法
	return c.ClientIP()
}
