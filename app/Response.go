package app

import (
	"github.com/fp/fp-gin-framework/app/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Request 上下文
type Request struct {
	Context *gin.Context
}

// Response 响应信息结构体
type Response struct {
	HttpCode int         `json:"-"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

// Option 定义配置选项函数（关键）
type Option func(*Response)

// C 设置JSON结构状态码
// @param int code 状态码
// @return Option 返回配置选项函数
func C(code int) Option {
	return func(this *Response) {
		this.Code = code
	}
}

// M 设置响应信息
// @param string message 自定义响应信息
// @return Option 返回配置选项函数
func M(message string) Option {
	return func(this *Response) {
		this.Message = message
	}
}

// D 设置响应参数
// @param interface{} data 响应数据
// @return Option 返回配置选项函数
func D(data interface{}) Option {
	return func(this *Response) {
		this.Data = data
	}
}

// H 设置HTTP响应状态码
// @param int HttpCode HTTP状态码，比如：200，500等
// @return Option 返回配置选项函数
func H(HttpCode int) Option {
	return func(this *Response) {
		this.HttpCode = HttpCode
	}
}

// Success 响应函数
// @param *Gin g 上下文结构体
// @param Option opts 接收多个配置选项函数参数，可以是C，M，D，H
func (r *Request) Success(opts ...Option) {
	defaultResponse := &Response{
		HttpCode: http.StatusOK,
		Code:     constants.SuccessCode,
		Message:  constants.Success,
		Data:     nil,
	}

	// 依次调用opts函数列表中的函数，为结构体成员赋值
	for _, o := range opts {
		o(defaultResponse)
	}

	r.Context.JSON(defaultResponse.HttpCode, defaultResponse)
	return
}
