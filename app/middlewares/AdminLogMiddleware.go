package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/constants/middleware"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/app/repositories"
	"net/http"
)

// AdminLogger 管理员操作日志
// @return gin.HandlerFunc
func AdminLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		var state int
		// 如果是登录请求则跳过
		if VerifyRoute(ctx.Request.URL.Path, ctx.Request.Method, map[string]string{"/admin/login": "POST"}) {

			return
		}
		admin, _ := ctx.Get("admin")
		info := admin.(*models.Admin)
		if ctx.Writer.Status() != http.StatusOK {
			state = 1
		} else {
			state = 2
		}
		_, err := repositories.AdminLog().Insert(map[string]interface{}{
			"type":    middleware.OperateType,
			"adminId": info.Id,
			"ip":      ctx.ClientIP(),
			"path":    ctx.Request.URL.Path,
			"method":  ctx.Request.Method,
			"code":    ctx.Writer.Status(),
			"state":   state,
		})
		if err != nil {
			panic(middleware.LogFail)
		}
	}
}
