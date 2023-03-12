package admin

const (
	NotExist              = 0
	UserNotExist          = "管理员不存在"
	PasswordError         = "密码错误"
	ErrorLoginOverload    = "登录次数过多，请30分钟后重试"
	GetAdminInfoFail      = "获取用户信息失败"
	LoginErrorLimitNumber = 3
	Increment             = 1
	KeyValidity           = 30
	AddFail               = "添加失败"
	RoleNotExist          = "角色不存在"
	AddRoleFail           = "添加角色失败"
	UpdateFail            = "更新失败"
	DeleteFail            = "删除角色失败"
	GetRoleFail           = "获取角色信息失败"
	DeleteAdminFail       = "删除失败"
	GetAdminListFail      = "获取列表失败"
	LoginSuccess          = "登录成功"
)
