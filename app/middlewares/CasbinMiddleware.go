package middlewares

import (
	"github.com/gin-gonic/gin"
	middlewareConstant "github.com/herman-hang/herman/app/constants/middleware"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/bootstrap/core"
)

// CheckPermission 权限检测
// @return gin.HandlerFunc
func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 合并两个map
		for k, v := range middlewareConstant.ExcludeRoute {
			middlewareConstant.ExcludeAuth[k] = v
		}
		if VerifyRoute(ctx.Request.URL.Path, ctx.Request.Method, middlewareConstant.ExcludeAuth) {
			return
		}
		admin, _ := ctx.Get("admin")
		info := admin.(*models.Admin)
		if info.Id == middlewareConstant.IsSuperAdmin {
			return
		}
		success, _ := core.Casbin.Enforce(info.User, ctx.Request.URL.Path, ctx.Request.Method)
		if !success {
			panic(middlewareConstant.PermissionDenied)
		}

		ctx.Next()
	}
}
