package controllers

import (
	"fp-back-user/app"
	"fp-back-user/app/constants"
	"github.com/gin-gonic/gin"
	"strings"
)

// GetParams 接收数据
// GET请求支持Query和Body接收数据
// POST只支持Body接收数据
func GetParams(ctx *gin.Context) (app.Gin, map[string]interface{}) {
	params := make(map[string]interface{})
	this := app.Gin{C: ctx}

	switch this.C.Request.Method {
	case "GET":
		if this.C.Request.URL.RawQuery != "" {
			for _, value := range strings.Split(this.C.Request.URL.RawQuery, "&") {
				paramSlice := strings.Split(value, "=")
				params[paramSlice[0]] = paramSlice[1]
			}
			break
		}

		if err := this.C.BindJSON(&params); err != nil {
			panic(err.Error())
		}
		break
	case "POST":
		if err := this.C.BindJSON(&params); err != nil {
			panic(err.Error())
		}
		break
	default:
		panic(constants.GetMessage(constants.MethodBan))
	}

	return this, params
}
