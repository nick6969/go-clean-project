package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/register"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Register(usecase *register.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析請求參數
		// 呼叫 usecase
		// 回傳結果
	}
}
