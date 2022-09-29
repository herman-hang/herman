package middlewares

import (
	"fp-back-user/app/utils"
	"github.com/gin-gonic/gin"
	"sort"
)

// 以下路由不校验token
var noVerify = []string{
	"/api/v1/user/login",
	"/api/v1/test",
}

// Jwt 鉴权
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if IsPath(noVerify, ctx.Request.URL.Path) {
			return
		}
		// 验证token
		utils.JwtVerify(ctx)
		ctx.Next()
	}
}

// IsPath 判断访问路径是否为白名单，在白名单直接返回true不验证token
func IsPath(strArray []string, target string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)

	// index的取值：0 ~ (len(str_array)-1)
	return index < len(strArray) && strArray[index] == target
}
