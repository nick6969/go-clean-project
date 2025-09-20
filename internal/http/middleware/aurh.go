package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/domain"
)

type AuthToken interface {
	ValidateAccessToken(token string) (int, *domain.GPError)
}

type Auth struct {
	token AuthToken
}

func NewAuth(token AuthToken) *Auth {
	return &Auth{
		token: token,
	}
}

func (a *Auth) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 從 Header 中取得 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(domain.NewGPError(domain.ErrCodeUnauthorized).Append("Authorization header is missing"))
			return
		}

		// 假設 token 格式為 "Bearer {token}"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Error(domain.NewGPError(domain.ErrCodeUnauthorized).Append("Authorization header format must be Bearer {token}"))
			return
		}

		token := parts[1]

		// 驗證 token
		userID, err := a.token.ValidateAccessToken(token)
		if err != nil {
			c.Error(domain.NewGPError(domain.ErrCodeUnauthorized).Append("Invalid token"))
			return
		}

		// 將 userID 存入 context，供後續處理使用
		c.Set("userID", userID)
	}
}
