package middlewares

import (
	"github.com/fp/fp-gin-framework/app"
	"github.com/fp/fp-gin-framework/app/constants"
	"github.com/gin-gonic/gin"
)

// CatchError 异常捕捉
// @return gin.HandlerFunc 返回一个中间件上下文
func CatchError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := &app.Request{Context: ctx}
		defer func() {
			if err := recover(); err != nil {
				// 没有定义
				response.Success(app.C(constants.ErrorCode), app.M(err.(string)))
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
