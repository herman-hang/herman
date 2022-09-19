package middlewares

import (
	"fp-back-user/app"
	"github.com/gin-gonic/gin"
)

// Jwt 鉴权
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		this := app.Gin{C: ctx}

		this.C.Next()
	}
}
