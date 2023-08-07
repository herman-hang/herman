package admin

import (
	"fmt"
	"github.com/herman-hang/herman/application/constants"
	CaptchaConstant "github.com/herman-hang/herman/application/constants/common/captcha"
	"github.com/herman-hang/herman/application/validates"
	"github.com/herman-hang/herman/kernel/app"
	utils2 "github.com/herman-hang/herman/kernel/utils"
	"github.com/mitchellh/mapstructure"
)

// CaptchaLoginValidate 管理员登录验证结构体
type CaptchaLoginValidate struct {
	User        string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password    string `json:"password" validate:"required,min=6,max=15" label:"密码"`
	CaptchaType int    `json:"captchaType" validate:"required,numeric,oneof=1 2" label:"验证码类型"`
	Token       string `json:"token" validate:"required" label:"验证码Token"`
	PointJson   string `json:"pointJson" validate:"required" label:"验证码PointJson"`
}

// ExcludeCaptchaLoginValidate 管理员登录排除验证码相关验证结构体
type ExcludeCaptchaLoginValidate struct {
	User     string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=15" label:"密码"`
}

// Login 登录验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func Login(data map[string]interface{}) (toMap map[string]interface{}) {
	// 判断是否需要验证码
	if !app.Config.Captcha.Switch {
		return excludeCaptchaLogin(data)
	}
	return captchaLogin(data)
}

// captchaLogin 验证码登录验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func captchaLogin(data map[string]interface{}) (toMap map[string]interface{}) {
	var login CaptchaLoginValidate
	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &login); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(login); err != nil {
		panic(err.Error())
	}

	// 验证码二次验证
	err := utils2.Factory().GetService(fmt.Sprintf("%s", data["captchaType"])).Verification(fmt.Sprintf("%s", data["token"]),
		fmt.Sprintf("%s", data["PointJson"]))
	if err != nil {
		panic(CaptchaConstant.CheckCaptchaError)
	}

	toMap, err = utils2.ToMap(&login, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}

// excludeCaptchaLogin 排除验证码登录验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func excludeCaptchaLogin(data map[string]interface{}) (toMap map[string]interface{}) {
	var login ExcludeCaptchaLoginValidate
	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &login); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(login); err != nil {
		panic(err.Error())
	}

	toMap, err := utils2.ToMap(&login, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}
