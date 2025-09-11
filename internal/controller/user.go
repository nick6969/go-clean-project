package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/controller/request"
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/register"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Register(usecase *register.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.Register
		// 1. 解析與驗證請求
		if err := c.ShouldBindJSON(&req); err != nil {
			// ShouldBindJSON 會自動根據 struct tag 進行驗證
			// 如果驗證失敗，會返回錯誤
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 2. 準備調用 UseCase 所需的資料
		input := register.NewInput(req.Email, req.Password)

		// 3. 調用 UseCase 的 Execute 方法
		output, err := usecase.Execute(c, input)

		// 4. 處理 UseCase 返回的結果
		if err != nil {
			// 專門處理已知的業務邏輯錯誤
			if errors.Is(err, errors.New("email is already registered")) {
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				return
			}

			// 對於其他未知的內部錯誤，回傳 500
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 5. 成功，回傳 200 狀態碼及 token
		c.JSON(http.StatusOK, gin.H{"data": output.AccessToken})
	}
}
