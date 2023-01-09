package middleware

// ExcludeRoute 以下路由不校验token
var ExcludeRoute = map[string]string{
	"/user/login": "post",
}

const (
	GuardError = "Guard错误"
)
