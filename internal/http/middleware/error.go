package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/domain"
	"github.com/nick6969/go-clean-project/internal/logger"
)

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h *ErrorHandler) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		loggerInstance, exists := c.Get("logger")
		if !exists {
			// 如果沒有 logger，無法記錄錯誤
			// 這會是程式邏輯錯誤，應該要確保每個請求都有 logger
			return
		}

		l, ok := loggerInstance.(logger.Logger)
		if !ok {
			// 如果 logger 類型不正確，無法記錄錯誤
			// 這會是程式邏輯錯誤，應該要確保 logger 類型正確
			return
		}

		// 取得最後一個錯誤
		err := c.Errors.Last().Err

		// 這裡假設 err 是 domain.GPError 類型
		if gpErr, ok := err.(domain.GPError); ok {
			if gpErr.HttpStatusCode() == http.StatusInternalServerError {
				l.Warn(c, "internal server error", gpErr)
			} else {
				l.Error(c, "client error", gpErr)
			}
			return
		}

		l.Error(c, "request error", err)
	}
}
