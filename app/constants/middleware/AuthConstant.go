package middleware

// ExcludeRoute 以下路由不校验token
var ExcludeRoute = map[string]string{
	"/user/login":  "post",
	"/admin/login": "post",
}

// ExcludeAuth 校验tokan，但不检查权限的路由
var ExcludeAuth = map[string]string{
	"/admin/index": "get",
}

const (
	GuardError = "Guard错误"
)
