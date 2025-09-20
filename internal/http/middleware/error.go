package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/domain"
	"github.com/nick6969/go-clean-project/internal/logger"
)

type ErrorHandler struct {
	logger logger.Logger
}

func NewErrorHandler(logger logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h *ErrorHandler) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		// 取得最後一個錯誤
		err := c.Errors.Last().Err

		// 這裡假設 err 是 domain.GPError 類型
		if gpErr, ok := err.(domain.GPError); ok {
			c.JSON(gpErr.HttpStatusCode(), gin.H{"error": gpErr.Message()})
			if gpErr.HttpStatusCode() == http.StatusInternalServerError {
				h.logger.Warn(c, "internal server error", gpErr)
			} else {
				h.logger.Error(c, "client error", gpErr)
			}
			return
		}

		h.logger.Error(c, "request error", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
