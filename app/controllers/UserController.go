package controllers

import (
	"fp-back-user/app"
	"fp-back-user/app/services"
	"github.com/gin-gonic/gin"
)

// UserList 用户列表
// controller中只负责接收数据，数据验证，函数调用，返回数据，不做其他业务处理，业务实现全部要在service中
func UserList(c *gin.Context) {
	// 必写步骤
	this := app.Gin{C: c}

	// 业务实现
	data := services.UserList()

	// Response参数可以设置一个或多个，也可以不设置
	this.Response(
		app.SetCode(200),
		app.SetMessage("操作成功"),
		app.SetData(data),
	)
}
