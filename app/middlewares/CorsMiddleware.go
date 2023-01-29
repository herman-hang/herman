package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
// @return gin.HandlerFunc
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, T, Sign")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, X-New-Token")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}
