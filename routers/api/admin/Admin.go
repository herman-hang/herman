package admin

import (
	"github.com/gin-gonic/gin"
	AdminController "github.com/herman-hang/herman/app/controllers/admin"
	MenuController "github.com/herman-hang/herman/app/controllers/menu"
	RoleController "github.com/herman-hang/herman/app/controllers/role"
)

// Router 后台相关路由
// @param *gin.RouterGroup router 路由组对象
// @return void
func Router(router *gin.RouterGroup) {
	// 管理员登录
	router.POST("/login", AdminController.Login)
	// 管理员添加
	router.POST("/admin", AdminController.AddAdmin)
	// 管理员修改
	router.PUT("/admin", AdminController.ModifyAdmin)
	// 根据ID查询管理员详情
	router.GET("/admin", AdminController.FindAdmin)
	// 管理员删除
	router.DELETE("/admin", AdminController.RemoveAdmin)
	// 管理员列表
	router.GET("/admin/list", AdminController.ListAdmin)

	// 添加角色
	router.POST("/role", RoleController.AddRole)
	// 删除角色
	router.DELETE("/role", RoleController.RemoveRole)
	// 修改角色
	router.PUT("/role", RoleController.ModifyRole)
	// 根据ID获取角色详情
	router.GET("/role", RoleController.FindRole)
	// 角色列表
	router.GET("/role/list", RoleController.ListRole)

	// 添加菜单
	router.POST("/menu", MenuController.AddMenu)
	// 修改菜单
	router.PUT("/menu", MenuController.ModifyMenu)
	// 根据ID获取菜单详情
	router.GET("/menu", MenuController.FindMenu)
	// 删除菜单
	router.DELETE("/menu", MenuController.RemoveMenu)
	// 菜单列表
	router.GET("/menu/list", MenuController.ListMenu)
}
