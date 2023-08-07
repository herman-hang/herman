package middlewares

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	"net/http"
	"runtime"
	"strings"
)

// CatchError 异常捕捉
// @return gin.HandlerFunc 返回一个中间件上下文
func CatchError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := application.Request{Context: ctx}
		defer func() {
			if data := recover(); data != nil {
				switch data.(type) {
				case string:
					data := fmt.Sprintf("%s", data)
					context.Json(nil, data, http.StatusInternalServerError)
				case map[string]interface{}:
					data := data.(map[string]interface{})
					context.Json(nil, data["message"], data["code"])
				default:
					fmt.Println(color.RedString(fmt.Sprintf("%s", data)))
					printStack()
				}
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

// printStack 获取当前的错误堆栈信息
// @return void
func printStack() {
	stack := make([]byte, 4096)
	length := runtime.Stack(stack, false)
	stackStr := string(stack[:length])
	stackLines := strings.Split(stackStr, "\n")
	for _, line := range stackLines {
		if !strings.Contains(line, "runtime/panic.go") {
			fmt.Println(color.RedString(line))
		}
	}
}
