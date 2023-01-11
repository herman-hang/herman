package middlewares

import "github.com/gin-gonic/gin"

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//user, _ := ctx.Get("user")

	}
}
