package app

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/constants"
	"io/ioutil"
)

// Request 请求结构体
type Request struct {
	Context *gin.Context
}

// Params 接收数据
// @param *gin.Context ctx 上下文
// @return params 返回一个二次封装上下文和响应对象
func (r *Request) Params() (params map[string]interface{}) {
	params = make(map[string]interface{})
	data, _ := r.Context.GetRawData()
	switch r.Context.Request.Method {
	case "GET":
		// query参数处理
		keys := r.Context.Request.URL.Query()
		if len(keys) != constants.LengthByZero {
			for k, v := range keys {
				params[k] = v[0]
			}
		} else {
			// body参数处理
			params = bodyParam(data, r.Context)
		}
	case "POST", "PUT", "DELETE":
		// body参数处理
		params = bodyParam(data, r.Context)
	default:
		panic(constants.MethodBan)
	}
	return params
}

// bodyParamHandle 统一处理Body参数
// @param []byte data 接收RawData
// @param *gin.Context ctx 上下文
// @return params 返回统一格式的数据
func bodyParam(data []byte, ctx *gin.Context) (params map[string]interface{}) {
	params = make(map[string]interface{})

	if len(data) != constants.LengthByZero {
		_ = json.Unmarshal(data, &params)
	}

	if len(params) == constants.LengthByZero {
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		_ = ctx.ShouldBindJSON(&params)
	}

	if len(params) == constants.LengthByZero {
		for k, v := range ctx.Request.Form {
			params[k] = v[0]
		}
	}

	return params
}
