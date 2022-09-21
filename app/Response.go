package app

import (
	"fp-back-user/app/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	HttpCode int         `json:"-"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

// Option 定义配置选项函数（关键）
type Option func(*Response)

// C 设置状态码
func C(code int) Option {
	return func(this *Response) {
		this.Code = code
	}
}

// M 设置响应信息
func M(message string) Option {
	return func(this *Response) {
		this.Message = message
	}
}

// D 设置响应参数
func D(data interface{}) Option {
	return func(this *Response) {
		this.Data = data
	}
}

// H 设置HTTP响应状态码
func H(HttpCode int) Option {
	return func(this *Response) {
		this.HttpCode = HttpCode
	}
}

// Response 响应函数
func (g *Gin) Response(opts ...Option) {

	defaultResponse := Response{
		HttpCode: http.StatusOK,
		Code:     constants.SUCCESS,
		Message:  constants.GetMessage(constants.SUCCESS),
		Data:     nil,
	}

	// 依次调用opts函数列表中的函数，为结构体成员赋值
	for _, o := range opts {
		o(&defaultResponse)
	}

	g.C.JSON(defaultResponse.HttpCode, defaultResponse)

	return
}
