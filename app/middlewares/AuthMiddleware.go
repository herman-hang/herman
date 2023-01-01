package middlewares

import (
	"github.com/fp/fp-gin-framework/app/constants"
	"github.com/fp/fp-gin-framework/app/repositories"
	"github.com/fp/fp-gin-framework/app/utils"
	"github.com/fp/fp-gin-framework/servers/settings"
	"github.com/gin-gonic/gin"
	"strings"
)

// 以下路由不校验token
var (
	prefix       = settings.Config.AppPrefix // 路由版本号
	excludeRoute = map[string]string{
		prefix + "/user/login": "post",
	}
)

// Jwt 鉴权
// @return gin.HandlerFunc 返回一个中间件上下文
func Jwt(guard string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if VerifyRoute(ctx.Request.URL.Path, ctx.Request.Method) {
			return
		}
		UserClaims := utils.JwtVerify(ctx, guard)
		switch guard {
		case "user", "mobile": // 前台和移动端（用户）
			// 用户信息存储在请求中
			ctx.Set("user", repositories.User.GetUserInfo(UserClaims.Uid))
		case "admin": // 中台（管理员）

		case "merchant": // 后台（商家）

		default:
			panic(constants.GuardError)
		}
		ctx.Next()
	}
}

// VerifyRoute 判断访问路径是否为白名单，在白名单直接返回true不验证token
// @param string route 当前请求路由
// @param string method 当前请求的http方法
// @return bool 返回一个路由是否存在不校验token数组路由中的值
func VerifyRoute(route string, method string) bool {
	if value, ok := excludeRoute[route]; !ok {
		return false
	} else if value == strings.ToUpper(method) {
		return true
	}
	return false
}
