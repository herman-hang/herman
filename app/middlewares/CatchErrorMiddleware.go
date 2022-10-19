package middlewares

import (
	"fp-back-user/app"
	"fp-back-user/app/constants"
	"github.com/gin-gonic/gin"
)

// CatchError 异常捕捉
// @return gin.HandlerFunc 返回一个中间件上下文
func CatchError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		this := app.Gin{C: ctx}
		defer func() {
			if err := recover(); err != nil {
				// 没有定义
				this.Response(app.C(constants.Error), app.M(err.(string)))
				this.C.Abort()
			}
		}()
		this.C.Next()
	}
}
