package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/common"
	middlewareConstant "github.com/herman-hang/herman/app/constants/middleware"
	"github.com/herman-hang/herman/app/models"
)

// CheckPermission 权限检测
// @return gin.HandlerFunc
func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		admin, _ := ctx.Get("admin")
		_ = admin.(*models.Admin)
		success, _ := common.Casbin.Enforce("admin", ctx.Request.URL.Path, ctx.Request.Method)
		if !success {
			panic(middlewareConstant.PermissionDenied)
		}
		ctx.Next()
	}
}
