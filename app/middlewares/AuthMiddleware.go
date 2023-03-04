package middlewares

import (
	"github.com/gin-gonic/gin"
	MiddlewareConstant "github.com/herman-hang/herman/app/constants/middleware"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/app/utils"
	"github.com/herman-hang/herman/servers/settings"
	"strings"
)

// Jwt 鉴权
// @return gin.HandlerFunc 返回一个中间件上下文
func Jwt(guard string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if VerifyRoute(ctx.Request.URL.Path, ctx.Request.Method, MiddlewareConstant.ExcludeRoute) {
			return
		}
		claims := utils.JwtVerify(ctx, guard)
		switch guard {
		case "user", "mobile": // 前台和移动端（用户）
			// 用户信息存储在请求中
			ctx.Set("user", repositories.User.GetUserInfo(claims.Uid))
		case "admin": // 管理员后台
			ctx.Set("admin", repositories.Admin.GetAdminInfo(claims.Uid))
		case "merchant": // 商家后台

		default:
			panic(MiddlewareConstant.GuardError)
		}
		ctx.Next()
	}
}

// VerifyRoute 判断访问路径是否为白名单，在白名单直接返回true不验证token
// @param string route 当前请求路由
// @param string method 当前请求的http方法
// @param map[string]string routes 路由组
// @return bool 返回一个路由是否存在不校验token数组路由中的值
func VerifyRoute(route string, method string, routes map[string]string) bool {
	attributes := make(map[string]string)
	for k, v := range routes {
		attributes[settings.Config.AppPrefix+k] = v
	}

	if value, ok := attributes[route]; !ok {
		return false
	} else if method == strings.ToUpper(value) {
		return true
	}
	return false
}
