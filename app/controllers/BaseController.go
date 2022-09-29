package controllers

import (
	"fp-back-user/app"
	"fp-back-user/app/constants"
	"github.com/gin-gonic/gin"
)

// GetParams 接收数据
// 目前只支持请求体批量接收数据
func GetParams(ctx *gin.Context) (app.Gin, map[string]interface{}) {
	var params map[string]interface{}

	this := app.Gin{C: ctx}
	switch this.C.Request.Method {
	case "GET":
		break
	case "POST":
		// 接收数据
		if err := this.C.BindJSON(&params); err != nil {
			panic(err.Error())
		}
		break
	case "PUT":
		break
	case "DELETE":
		break
	default:
		panic(constants.GetMessage(constants.MethodBan))
	}

	return this, params
}
