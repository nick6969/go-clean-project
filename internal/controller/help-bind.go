package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nick6969/go-clean-project/internal/domain"
)

// jsonResponseHandler handles JSON requests and responses.
// tag is `json:"xxx"`
func bindJSON[Request any](ctx *gin.Context) (Request, bool) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiFailWithErrorCode(ctx, domain.ErrCodeParametersNotCorrect, err)
		return req, false
	}
	return req, true
}

// // bindQuery binds the query parameters from the request context to the provided struct.
// // tag is `form:"xxxx"`
// func bindQuery[Request any](ctx *gin.Context) (Request, bool) {
// 	var req Request
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		apiFailWithErrorCode(ctx, domain.ErrCodeParametersNotCorrect, err)
// 		return req, false
// 	}
// 	return req, true
// }

// // bindURI binds the query parameters from the route, e.g. "/:id"
// // tag is `uri:"xxxx"`
// func bindURI[Request any](ctx *gin.Context) (Request, bool) {
// 	var req Request
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		apiFailWithErrorCode(ctx, domain.ErrCodeParametersNotCorrect, err)
// 		return req, false
// 	}
// 	return req, true
// }
