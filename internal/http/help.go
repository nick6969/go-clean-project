package http

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func staticFile(r *gin.RouterGroup, fs embed.FS, path, filePath, contentType string, cacheTime int) {
	r.GET(path, func(ctx *gin.Context) {
		file, err := fs.ReadFile(filePath)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		if cacheTime > 0 {
			ctx.Header("Cache-Control", fmt.Sprintf("max-age=%d", cacheTime))
		} else {
			ctx.Header("Cache-Control", "no-cache")
		}
		ctx.Data(http.StatusOK, contentType, file)
	})
}
