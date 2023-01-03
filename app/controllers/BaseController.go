package controllers

import (
	"encoding/json"
	"github.com/fp/fp-gin-framework/app"
	"github.com/fp/fp-gin-framework/app/constants"
	"github.com/gin-gonic/gin"
)

// GetParams 接收数据
// @return this params 返回一个二次封装上下文和接收前端数据参数
func GetParams(ctx *gin.Context) (params map[string]interface{}, response app.Request) {
	params = make(map[string]interface{})
	response = app.Request{Context: ctx}
	data, _ := ctx.GetRawData()
	switch ctx.Request.Method {
	case "GET":
		// query参数处理
		keys := ctx.Request.URL.Query()
		if len(keys) != 0 {
			for k, v := range keys {
				params[k] = v[0]
			}
		}

		// body参数处理
		if len(data) != 0 {
			_ = json.Unmarshal(data, &params)
		}
	case "POST", "PUT", "DELETE":
		if len(string(data)) != 0 {
			_ = json.Unmarshal(data, &params)
		}
	default:
		panic(constants.MethodBan)
	}

	return params, response
}
