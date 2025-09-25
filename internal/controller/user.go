package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/controller/request"
	"github.com/nick6969/go-clean-project/internal/domain"
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/changePassword"
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/login"
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/register"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// @Summary			Register
// @Description	User Registration
// @Tags				User
// @Accept			json
// @Produce			json
// @Param				request	body		request.Register	true	"User Registration Info"
// @Success			200			{object}	register.Output
// @Failure			400			{object}	generalErrorResponse
// @Failure			500			{object}	generalErrorResponse
// @Router			/api/user/register [post]
func (u *UserController) Register(useCase *register.UseCase) gin.HandlerFunc {
	return jsonResponseHandler(func(ctx *gin.Context, req request.Register) (*register.Output, *domain.GPError) {
		input := register.NewInput(req.Email, req.Password)
		return useCase.Execute(ctx, input)
	})
}

// @Summary			Login
// @Description	User Login
// @Tags				User
// @Accept			json
// @Produce			json
// @Param				request	body		request.Login	true	"User Login Info"
// @Success			200			{object}	login.Output
// @Failure			400			{object}	generalErrorResponse
// @Failure			500			{object}	generalErrorResponse
// @Router			/api/user/login [post]
func (u *UserController) Login(useCase *login.UseCase) gin.HandlerFunc {
	return jsonResponseHandler(func(ctx *gin.Context, req request.Login) (*login.Output, *domain.GPError) {
		input := login.NewInput(req.Email, req.Password)
		return useCase.Execute(ctx, input)
	})
}

// @Summary			Change Password
// @Description	Change User Password
// @Tags				User
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param				request	body		request.ChangePassword	true	"Change Password Info"
// @Success			204
// @Failure			400			{object}	generalErrorResponse
// @Failure			500			{object}	generalErrorResponse
// @Router			/api/user/change-password [post]
func (u *UserController) ChangePassword(useCase *changePassword.UseCase) gin.HandlerFunc {
	return jsonUserIDHandler(func(ctx *gin.Context, userID int, req request.ChangePassword) *domain.GPError {
		input := changePassword.NewInput(userID, req.Password, req.NewPassword)
		return useCase.Execute(ctx, input)
	})
}
