package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	"github.com/herman-hang/herman/app/middlewares"
	"github.com/herman-hang/herman/routers"
	"github.com/herman-hang/herman/servers"
	"github.com/herman-hang/herman/servers/settings"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCase 测试用例
type TestCase struct {
	Method   string                 // 请求方法
	Uri      string                 // 链接
	Params   map[string]interface{} // 参数
	Code     int                    // 返回状态码
	Message  string                 // 返回信息
	Desc     string                 // 描述
	ShowBody bool                   // 是否展示返回
}

// Call 测试基类方法
// @param *testing.T t 测试对象
// @param []TestCase testcase 测试用例切片
// @return void
func Call(t *testing.T, testcase []TestCase) {
	setup()
	for k, v := range testcase {
		var response app.Response
		// map转json
		params, err := json.Marshal(v.Params)
		assert.NoError(t, err)
		_, _, w := request(v.Method, v.Uri, string(params))
		fmt.Printf("第%d个测试用例----接口:%s URI:%s ", k+1, v.Desc, v.Uri)
		if v.ShowBody {
			fmt.Printf("响应数据:%s\n", w.Body.String())
		}
		// json转struct
		err = json.Unmarshal([]byte(w.Body.String()), &response)
		assert.NoError(t, err)
		// json转map
		assert.Equal(t, v.Code, response.Code, "响应自定义状态码不一致")
		assert.Equal(t, v.Message, response.Message, "响应自定义信息不一致")
	}
}

// setup 开始测试前初始化
// @return void
func setup() {
	settings.InitConfig()
	servers.ZapLogs()
	middlewares.Reload()
}

// request 发起请求
// @param string method 请求方法
// @param string url 请求链接
// @param string body 请求参数
func request(method string, uri string, body string) (
	c *gin.Context,
	r *http.Request,
	w *httptest.ResponseRecorder,
) {
	gin.SetMode(gin.TestMode)
	e := gin.Default()
	e = routers.InitRouter(e)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	r = httptest.NewRequest(method, uri, bytes.NewBufferString(body))
	c.Request = r
	c.Request.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, r)
	return
}
