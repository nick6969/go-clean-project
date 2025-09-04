package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/application"
)

func registerRoutes(engine *gin.Engine, app *application.Application) {
	registerRootRoutes(engine, app)
	registerAPIRoutes(engine.Group("api"), app)
}
