package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("admin")
		fmt.Println(user)
		/*		cachedEnforcer := utils.Enforcer(utils.GetAdminPolicy())
				success, _ := cachedEnforcer.Enforce(user.role, ctx.Request.URL.Path, ctx.Request.Method)
				if !success {
					panic(middlewareConstants.PermissionDenied)
				}
				ctx.Next()*/
	}
}
