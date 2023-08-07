package middleware

// ExcludeRoute 以下路由不校验token
var ExcludeRoute = map[string]string{
	"/user/login":         "POST",
	"/admin/login":        "POST",
	"/pc/login":           "POST",
	"/pc/send/code":       "POST",
	"/pc/chat/gpt/stream": "GET",
}

// ExcludeAuth 校验tokan，但不检查权限的路由
var ExcludeAuth = map[string]string{
	"/admin/index": "GET",
}

const (
	GuardError = "Guard错误"
)
