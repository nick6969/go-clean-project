package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/application"
)

func registerRootRoutes(r gin.IRouter, _ *application.Application) {
	// Register your API routes here
	r.GET("health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
