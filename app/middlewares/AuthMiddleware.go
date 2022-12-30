package middlewares

import (
	"github.com/fp/fp-gin-framework/app/repositories"
	"github.com/fp/fp-gin-framework/app/utils"
	"github.com/gin-gonic/gin"
	"sort"
)

// 以下路由不校验token
var noVerify = []string{
	"/api/v1/user/login",
}

// Jwt 鉴权
// @return gin.HandlerFunc 返回一个中间件上下文
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if IsPath(noVerify, ctx.Request.URL.Path) {
			return
		}
		UserClaims := utils.JwtVerify(ctx)
		// 用户信息存储在请求中
		ctx.Set("user", repositories.UserInfo(UserClaims.Uid))
		ctx.Next()
	}
}

// IsPath 判断访问路径是否为白名单，在白名单直接返回true不验证token
// @param []string strArray 路由不校验token数组
// @param string target 待校验的路由
// @return bool 返回一个路由是否存在不校验token数组路由中的值
func IsPath(strArray []string, target string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)

	// index的取值：0 ~ (len(str_array)-1)
	return index < len(strArray) && strArray[index] == target
}
