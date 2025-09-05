package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/application"
	"github.com/nick6969/go-clean-project/internal/controller"
)

func registerAPIRoutes(r gin.IRouter, app *application.Application) {
	registerAPIUserRoutes(r.Group("user"), app)
}

func registerAPIUserRoutes(r gin.IRouter, app *application.Application) {
	uc := controller.NewUserController()
	r.POST("register", uc.Register(app.UseCase.User.Register))
}
