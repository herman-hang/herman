package middlewares

import (
	middlewareConstants "github.com/fp/fp-gin-framework/app/constants/middleware"
	"github.com/fp/fp-gin-framework/app/models"
	"github.com/fp/fp-gin-framework/app/utils"
	"github.com/gin-gonic/gin"
)

// CheckPermission 权限检测
// @return gin.HandlerFunc
func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		admin, _ := ctx.Get("admin")
		waitUse := admin.(*models.Admin)
		cachedEnforcer := utils.Enforcer(utils.GetAdminPolicy())
		success, _ := cachedEnforcer.Enforce(waitUse.Role, ctx.Request.URL.Path, ctx.Request.Method)
		if !success {
			panic(middlewareConstants.PermissionDenied)
		}
		ctx.Next()
	}
}
