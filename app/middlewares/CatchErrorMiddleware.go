package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	"net/http"
)

// CatchError 异常捕捉
// @return gin.HandlerFunc 返回一个中间件上下文
func CatchError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := app.Request{Context: ctx}
		defer func() {
			if data := recover(); data != nil {
				switch data.(type) {
				case string:
					context.Json(nil, fmt.Sprintf("%s", data), http.StatusInternalServerError)
				case map[string]interface{}:
					data := data.(map[string]interface{})
					context.Json(nil, data["message"], data["code"])
				}
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
