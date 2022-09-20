package middlewares

import (
	"fp-back-user/app"
	"fp-back-user/app/utils"
	"github.com/gin-gonic/gin"
)

// Jwt 鉴权
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		this := app.Gin{C: ctx}
		utils.JwtVerify(ctx)
		this.C.Next()
	}
}
