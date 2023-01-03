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
var excludeRoute = map[string]string{
	"/user/login": "post",
}

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
		case "admin": // 管理员后台

		case "merchant": // 商家后台

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
	attributes := make(map[string]string)
	for k, v := range excludeRoute {
		attributes[settings.Config.AppPrefix+k] = v
	}

	if value, ok := attributes[route]; !ok {
		return false
	} else if method == strings.ToUpper(value) {
		return true
	}
	return false
}
