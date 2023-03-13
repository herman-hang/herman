package test

import (
	"bytes"
	"encoding/json"
	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	MiddlewareConstant "github.com/herman-hang/herman/app/constants/middleware"
	"github.com/herman-hang/herman/app/middlewares"
	"github.com/herman-hang/herman/bootstrap/core"
	"github.com/herman-hang/herman/routers"
	"github.com/herman-hang/herman/servers"
	"github.com/herman-hang/herman/servers/settings"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
)

// SuiteCase 测试套件
type SuiteCase struct {
	suite.Suite
	Guard         string
	Authorization string
	AppPrefix     string
}

// Case 测试用例
type Case struct {
	Method  string                 // 请求方法
	Uri     string                 // 请求路由
	Params  map[string]interface{} // 请求参数
	Code    int                    // 响应自定义状态码
	Message string                 // 响应自定义信息
	IsList  bool                   // 是否是列表
	Fields  []string               // 断言字段
}

// SetupSuite 测试套件前置函数
// @return void
func (s *SuiteCase) SetupSuite() {
	settings.InitConfig()
	servers.ZapLogs()
	middlewares.Reload()
	gin.SetMode(settings.Config.Mode)
	e := gin.Default()
	e.Use(middlewares.CatchError())
	core.Engine = routers.InitRouter(e)
	s.AppPrefix = settings.Config.AppPrefix
	switch s.Guard {
	case "admin":
		s.AdminLogin()
	default:
		panic(MiddlewareConstant.GuardError)
	}
}

// Assert 断言
// @param []Case testCase 测试用例切片
// @return void
func (s *SuiteCase) Assert(testCase []Case) {
	var response app.Response
	for _, v := range testCase {
		_, _, w := s.Request(v.Method, v.Uri, v.Params)
		// json转struct
		err := json.Unmarshal(w.Body.Bytes(), &response)
		s.T().Logf("Response: %s", w.Body.String())
		assert.Equal(s.T(), err, nil)
		assert.Equal(s.T(), v.Code, response.Code)
		assert.Equal(s.T(), v.Message, response.Message)
		// 是否为列表
		if v.IsList {
			for _, field := range v.Fields {
				data := response.Data.(map[string]interface{})["list"].([]interface{})
				for _, datum := range data {
					// 不相等测试通过
					assert.NotEqual(s.T(), nil, datum.(map[string]interface{})[field])
				}
			}
		} else {
			switch response.Data.(type) {
			case map[string]interface{}: // 非数组
				for _, field := range v.Fields {
					// 不相等测试通过
					assert.NotEqual(s.T(), nil, response.Data.(map[string]interface{})[field])
				}
			case []interface{}: // 数组
				for _, field := range v.Fields {
					for _, datum := range response.Data.([]interface{}) {
						// 不相等测试通过
						assert.NotEqual(s.T(), nil, datum.(map[string]interface{})[field])
					}
				}
			}
		}
	}
}

// Request 发起请求
// @param string method 请求方法
// @param string url 请求链接
// @param map[string]interface{} body 请求参数
func (s *SuiteCase) Request(method string, uri string, body map[string]interface{}) (
	c *gin.Context,
	r *http.Request,
	w *httptest.ResponseRecorder,
) {
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	// map转json
	requestBody, _ := json.Marshal(body)
	r = httptest.NewRequest(method, uri, bytes.NewBuffer(requestBody))
	r.Header.Set("Content-Type", "application/json")
	if len(s.Authorization) > 0 {
		r.Header.Set("Authorization", s.Authorization)
	}
	c.Request = r
	core.Engine.ServeHTTP(w, r)
	return
}

// AdminLogin 管理员登录
// @return void
func (s *SuiteCase) AdminLogin() {
	var response app.Response
	// map转json
	_, _, w := s.Request("POST", "/api/v1/admin/login", map[string]interface{}{
		"user":     "admin",
		"password": "123456",
	})
	// json转struct
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	s.Authorization = response.Data.(string)
}

// TearDownSuite 测试套件后置函数
// @return void
func (s *SuiteCase) TearDownSuite() {
	middlewares.Close()
}
