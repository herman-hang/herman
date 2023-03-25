package user

const (
	TokenNotExit          = "Token不存在"
	TokenNotValid         = "Token无效"
	CreateTokenFail       = "创建Token失败"
	TokenExpires          = "登录超时"
	PasswordError         = "密码错误"
	TokenError            = "请求头中Authorization格式有误"
	TokenRefreshFail      = "Token刷新失败"
	TokenSaveFail         = "Token保存失败"
	GetUserInfoFail       = "获取用户信息失败"
	ErrorLoginOverload    = "登录次数过多，请30分钟后重试"
	LoginErrorLimitNumber = 3
	Increment             = 1
	KeyValidity           = 30
	LengthByZero          = 0
	SplitByTwo            = 2
	Minute                = 10
)
