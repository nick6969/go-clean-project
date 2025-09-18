package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthToken interface {
	ValidateAccessToken(token string) (int, error)
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// 假設 token 格式為 "Bearer {token}"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		token := parts[1]

		// 驗證 token
		userID, err := a.token.ValidateAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// 將 userID 存入 context，供後續處理使用
		c.Set("userID", userID)
	}
}
