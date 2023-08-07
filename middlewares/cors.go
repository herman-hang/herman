package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/kernel/app"
	"net/http"
	"strings"
)

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
// @return gin.HandlerFunc
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查请求的Origin是否在允许的列表中
		var isAllowed bool
		allowedOrigins := strings.Split(app.Config.Cores.Origins, ",")
		origin := ctx.GetHeader("Origin")

		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, T, Sign")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, X-New-Token")
			ctx.Header("Access-Control-Allow-Credentials", "true")

			// 放行所有OPTIONS方法
			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(http.StatusNoContent)
				return
			}
		}

		// 处理请求
		ctx.Next()
	}
}
