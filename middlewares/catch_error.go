package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	"github.com/herman-hang/herman/kernel/core"
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
					data := fmt.Sprintf("%s", data)
					core.Log.Errorln(data)
					context.Json(nil, data, http.StatusInternalServerError)
				case map[string]interface{}:
					data := data.(map[string]interface{})
					core.Log.Errorln(data)
					context.Json(nil, data["message"], data["code"])
				}
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
