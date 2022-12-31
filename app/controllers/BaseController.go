package controllers

import (
	"bytes"
	"github.com/fp/fp-gin-framework/app"
	"github.com/fp/fp-gin-framework/app/constants"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

// GetParams 接收数据
// @param *gin.Context ctx 上下文
// @return this params 返回一个二次封装上下文和接收前端数据参数
func GetParams(ctx *gin.Context) (this app.Gin, params map[string]interface{}) {
	params = make(map[string]interface{})
	this = app.Gin{C: ctx}

	data, _ := this.C.GetRawData()
	// GET请求支持Query和Body接收数据
	// POST只支持Body接收数据
	switch this.C.Request.Method {
	case "GET":
		if this.C.Request.URL.RawQuery != "" {
			for _, value := range strings.Split(this.C.Request.URL.RawQuery, "&") {
				paramSlice := strings.Split(value, "=")
				params[paramSlice[0]] = paramSlice[1]
			}
		} else if string(data) != "" {
			this.C.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			if err := this.C.ShouldBindJSON(&params); err != nil {
				panic(err.Error())
			}
		}
	case "POST", "PUT", "DELETE":
		if string(data) != "" {
			this.C.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			if err := this.C.ShouldBindJSON(&params); err != nil {
				panic(err.Error())
			}
		}
	default:
		panic(constants.MethodBan)
	}

	return this, params
}
