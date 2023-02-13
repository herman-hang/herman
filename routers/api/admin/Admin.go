package admin

import (
	"github.com/gin-gonic/gin"
	AdminController "github.com/herman/app/controllers/admin"
	RoleController "github.com/herman/app/controllers/role"
	"github.com/herman/app/middlewares"
)

// Router 后台相关路由
// @param *gin.RouterGroup router 路由组对象
// @return void
func Router(router *gin.RouterGroup) {
	router.Use(middlewares.Jwt("admin"))
	{
		// 管理员登录
		router.POST("/login", AdminController.Login)

		// 添加角色
		router.POST("/role", RoleController.AddRole)
		// 删除角色
		router.DELETE("/role", RoleController.DeleteRole)
		// 修改角色
		router.PUT("/role", RoleController.ModifyRole)
		// 根据ID获取角色详情
		router.GET("/role", RoleController.FindRole)
		// 角色列表
		router.GET("/roles", RoleController.ListRole)
	}
}
