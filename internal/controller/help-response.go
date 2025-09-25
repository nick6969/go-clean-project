package controller

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/domain"
)

type generalSuccessResponse[T any] struct {
	// Response data
	Data T `json:"data"`
}

type generalErrorResponse struct {
	// Error code
	Code int `json:"code" example:"400000"`
	// Error message
	Message string `json:"message" example:"Parameters Not Correct"`
}

// Success Area
func apiSuccess(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNoContent)
}

func apiSuccessWithData[T any](ctx *gin.Context, data T) {
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		generalSuccessResponse[T]{Data: data},
	)
}

func apiSuccessWithObject(ctx *gin.Context, data any) {
	ctx.AbortWithStatusJSON(http.StatusOK, data)
}

// Failed Area
func apiFailWithErrorCode(ctx *gin.Context, code domain.GPErrorCode, err error) {
	glError := domain.NewGPErrorWithError(code, err)
	apiFailWithGPError(ctx, glError)
}

func apiFailWithGPError(ctx *gin.Context, err *domain.GPError) {
	ctx.JSON(err.HttpStatusCode(), generalErrorResponse{
		Code:    int(err.ErrorCode()),
		Message: err.Message(),
	})
	// 把錯誤記錄到 Gin 的 context 中，讓 ErrorHandler middleware 可以記錄
	ctx.Error(err)
}

func handleResponse[Response any](ctx *gin.Context, response Response, err *domain.GPError) {
	if err != nil {
		apiFailWithGPError(ctx, err)
		return
	}

	kind := reflect.TypeOf(response).Kind()
	if kind == reflect.Array || kind == reflect.Slice || kind == reflect.String || kind == reflect.Bool {
		apiSuccessWithData(ctx, response)
	} else {
		apiSuccessWithObject(ctx, response)
	}
}

func handleError(ctx *gin.Context, err *domain.GPError) {
	if err != nil {
		apiFailWithGPError(ctx, err)
	} else {
		apiSuccess(ctx)
	}
}
