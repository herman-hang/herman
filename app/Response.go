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
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Option 定义配置选项函数（关键）
type Option func(*Response)

// SetCode 设置状态码
func SetCode(code int) Option {
	return func(this *Response) {
		this.Code = code
	}
}

// SetMessage 设置响应信息
func SetMessage(message string) Option {
	return func(this *Response) {
		this.Message = message
	}
}

// SetData 设置响应参数
func SetData(data interface{}) Option {
	return func(this *Response) {
		this.Data = data
	}
}

// Response 响应函数
func (g *Gin) Response(opts ...Option) {

	defaultResponse := Response{
		Code:    constants.SUCCESS,
		Message: constants.GetMessage(constants.SUCCESS),
		Data:    nil,
	}

	// 依次调用opts函数列表中的函数，为结构体成员赋值
	for _, o := range opts {
		o(&defaultResponse)
	}

	g.C.JSON(http.StatusOK, defaultResponse)
	return
}
