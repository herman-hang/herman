package admin

import (
	"github.com/gin-gonic/gin"
	AdminController "github.com/herman-hang/herman/app/controllers/admin"
	DictionaryController "github.com/herman-hang/herman/app/controllers/dictionary"
	FileController "github.com/herman-hang/herman/app/controllers/file"
	MenuController "github.com/herman-hang/herman/app/controllers/menu"
	RoleController "github.com/herman-hang/herman/app/controllers/role"
	SystemController "github.com/herman-hang/herman/app/controllers/system"
)

// Router 后台相关路由
// @param *gin.RouterGroup router 路由组对象
// @return void
func Router(router *gin.RouterGroup) {
	// 管理员登录
	router.POST("/login", AdminController.Login)
	// 管理员添加
	router.POST("/admins", AdminController.AddAdmin)
	// 管理员修改
	router.PUT("/admins", AdminController.ModifyAdmin)
	// 根据ID查询管理员详情
	router.GET("/admins/:id", AdminController.FindAdmin)
	// 管理员删除
	router.DELETE("/admins", AdminController.RemoveAdmin)
	// 管理员列表
	router.GET("/admins", AdminController.ListAdmin)
	// 管理员日志列表
	router.GET("/admin/logs", AdminController.LogList)

	// 添加角色
	router.POST("/roles", RoleController.AddRole)
	// 删除角色
	router.DELETE("/roles", RoleController.RemoveRole)
	// 修改角色
	router.PUT("/roles", RoleController.ModifyRole)
	// 根据ID获取角色详情
	router.GET("/roles/:id", RoleController.FindRole)
	// 角色列表
	router.GET("/roles", RoleController.ListRole)

	// 添加菜单
	router.POST("/menus", MenuController.AddMenu)
	// 修改菜单
	router.PUT("/menus", MenuController.ModifyMenu)
	// 根据ID获取菜单详情
	router.GET("/menus/:id", MenuController.FindMenu)
	// 删除菜单
	router.DELETE("/menus", MenuController.RemoveMenu)
	// 菜单列表
	router.GET("/menus", MenuController.ListMenu)

	// 添加数据字典
	router.POST("/dictionaries", DictionaryController.AddDictionary)
	// 删除数据字典
	router.DELETE("/dictionaries", DictionaryController.RemoveDictionary)
	// 修改数据字典
	router.PUT("/dictionaries", DictionaryController.ModifyDictionary)
	// 根据数据字典KEY获取明细值
	router.GET("/dictionaries/details", DictionaryController.DetailsByDictionary)
	// 根据ID获取数据字典详情
	router.GET("/dictionaries/:id", DictionaryController.FindDictionary)
	// 数据字典列表
	router.GET("/dictionaries", DictionaryController.ListDictionary)

	// 添加明细值
	router.POST("/dictionaries/details", DictionaryController.AddDetail)
	// 删除明细值
	router.DELETE("/dictionaries/details", DictionaryController.RemoveDetail)
	// 根据ID获取明细值详情
	router.GET("/dictionaries/details/:id", DictionaryController.FindDetail)
	// 修改明细值
	router.PUT("/dictionaries/details", DictionaryController.ModifyDetail)

	// 文件上传
	router.POST("/files/uploads", FileController.UploadFile)
	// 文件下载
	router.GET("/files/download/:id", FileController.DownloadFile)
	// 图片预览
	router.GET("/files/preview/:id", FileController.PreviewFile)

	// 系统设置信息
	router.GET("/system", SystemController.FindSystem)
	// 修改系统设置信息
	router.PUT("/system", SystemController.ModifySystem)

}
