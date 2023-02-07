package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/herman/app"
	"github.com/herman/app/constants"
	"io/ioutil"
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
		if len(keys) != constants.LengthByZero {
			for k, v := range keys {
				params[k] = v[0]
			}
		} else {
			// body参数处理
			params = bodyParamHandle(data, ctx)
		}
	case "POST", "PUT", "DELETE":
		// body参数处理
		params = bodyParamHandle(data, ctx)
	default:
		panic(constants.MethodBan)
	}

	return params, response
}

// bodyParamHandle 统一处理Body参数
// @param []byte data 接收RawData
// @param *gin.Context ctx 上下文
// @return params 返回统一格式的数据
func bodyParamHandle(data []byte, ctx *gin.Context) (params map[string]interface{}) {
	params = make(map[string]interface{})
	if len(string(data)) != constants.LengthByZero {
		_ = json.Unmarshal(data, &params)
	}

	if len(params) == constants.LengthByZero {
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		_ = ctx.ShouldBind(&params)
	}

	if len(params) == constants.LengthByZero {
		for k, v := range ctx.Request.Form {
			params[k] = v[0]
		}
	}

	return params
}
