package controllers

import (
	"fp-back-user/app"
	"fp-back-user/app/common"
	"fp-back-user/app/constants"
	"fp-back-user/app/services"
	"github.com/gin-gonic/gin"
)

// UserLogin 用户列表
// controller中只负责接收数据，数据验证，函数调用，返回数据，不做其他业务处理，业务实现全部要在service中
func UserLogin(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	this, data := common.GetParams(ctx)

	// 验证数据

	// Response参数可以设置一个或多个，也可以不设置
	this.Response(
		app.C(constants.SUCCESS),
		app.M(constants.GetMessage(constants.SUCCESS)),
		app.D(services.UserLogin(data)),
	)
	return
}
