package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application/constants/admin/middleware"
	"github.com/herman-hang/herman/application/models"
	"github.com/herman-hang/herman/kernel/core"
)

// CheckPermission 权限检测
// @return gin.HandlerFunc
func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 合并两个map
		for k, v := range middleware.ExcludeRoute {
			middleware.ExcludeAuth[k] = v
		}
		if VerifyRoute(ctx.Request.URL.Path, ctx.Request.Method, middleware.ExcludeAuth) {
			return
		}
		admin, _ := ctx.Get("admin")
		info := admin.(*models.Admin)
		if info.Id == middleware.IsSuperAdmin {
			return
		}
		success, _ := core.Casbin().Enforce(info.User, ctx.Request.URL.Path, ctx.Request.Method)
		if !success {
			panic(middleware.PermissionDenied)
		}

		ctx.Next()
	}
}
