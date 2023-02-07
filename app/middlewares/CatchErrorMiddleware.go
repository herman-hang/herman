package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman/app"
	"github.com/herman/app/common"
	"net/http"
)

// CatchError 异常捕捉
// @return gin.HandlerFunc 返回一个中间件上下文
func CatchError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := &app.Request{Context: ctx}
		defer func() {
			if err := recover(); err != nil {
				errorMessage := fmt.Sprintf("%s", err)
				// 日志记录
				common.Log.Errorf(errorMessage)
				// 没有定义
				response.Success(app.C(http.StatusInternalServerError), app.M(errorMessage))
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
