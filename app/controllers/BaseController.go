package controllers

import (
	"fp-back-user/app"
	"github.com/gin-gonic/gin"
)

// GetParams 接收数据
// 目前只支持请求体批量接收数据
func GetParams(ctx *gin.Context) (app.Gin, map[string]interface{}) {

	var params map[string]interface{}

	this := app.Gin{C: ctx}
	// 接收数据
	if err := this.C.BindJSON(&params); err != nil {
		panic(err)
	}

	return this, params
}
