package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/domain"
)

func jsonResponseHandler[Request, Response any](f func(*gin.Context, Request) (Response, *domain.GPError)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, ok := bindJSON[Request](ctx)
		if !ok {
			return
		}
		resp, err := f(ctx, req)
		handleResponse(ctx, resp, err)
	}
}

func jsonUserIDHandler[Request any](f func(*gin.Context, int, Request) *domain.GPError) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("userID")
		if userID == 0 {
			apiFailWithErrorCode(ctx, domain.ErrCodeUnauthorized, errors.New("unauthorized"))
			return
		}

		req, ok := bindJSON[Request](ctx)
		if !ok {
			return
		}

		err := f(ctx, userID, req)
		handleError(ctx, err)
	}
}
