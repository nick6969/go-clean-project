package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of http requests.",
		},
		[]string{"method", "path", "status"}, // 標籤
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of http requests.",
		},
		[]string{"method", "path"},
	)
)

type Metrics struct {
	// Placeholder for future metrics fields
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		path := c.FullPath() // 使用路由範本，避免維度爆炸
		method := c.Request.Method
		status := strconv.Itoa(c.Writer.Status())

		// 更新指標
		httpRequestDuration.WithLabelValues(method, path).Observe(duration.Seconds())
		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
	}
}
