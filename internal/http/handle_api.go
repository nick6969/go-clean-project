package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/application"
)

func registerAPIRoutes(r gin.IRouter, _ *application.Application) {
	// Register your API routes here
	r.GET("hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
