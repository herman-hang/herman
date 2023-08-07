package utils

import (
	"fmt"
	CaptchaConfig "github.com/TestsLing/aj-captcha-go/config"
	Constant "github.com/TestsLing/aj-captcha-go/const"
	CaptchaService "github.com/TestsLing/aj-captcha-go/service"
	CaptchaConstant "github.com/herman-hang/herman/application/constants/common/captcha"
	"github.com/herman-hang/herman/kernel/app"
	"math/rand"
	"strconv"
	"time"
)

// Factory 初始化滑块验证码
// @return factory 返回一个验证码工厂
func Factory() (factory *CaptchaService.CaptchaServiceFactory) { // 行为校验配置模块（具体参数可从业务系统配置文件自定义）
	// 行为校验初始化
	factory = CaptchaService.NewCaptchaServiceFactory(
		CaptchaConfig.BuildConfig(app.Config.Captcha.CacheType,
			app.Config.Captcha.ResourcePath,
			&CaptchaConfig.WatermarkConfig{
				Text: app.Config.Captcha.Text,
			},
			nil, nil, app.Config.Captcha.CacheExpireSec))
	// 注册内存缓存
	factory.RegisterCache(Constant.MemCacheKey, CaptchaService.NewMemCacheService(CaptchaConstant.CacheMaxNumber))
	// 注册自定义配置redis数据库
	factory.RegisterCache(Constant.RedisCacheKey, CaptchaService.NewConfigRedisCacheService([]string{fmt.Sprintf("%s:%d",
		app.Config.Redis.Host,
		app.Config.Redis.Port,
	)},
		app.Config.Redis.UserName,
		app.Config.Redis.Password,
		false,
		app.Config.Redis.Db,
	))
	// 注册文字点选验证码服务
	factory.RegisterService(Constant.ClickWordCaptcha, CaptchaService.NewClickWordCaptchaService(factory))
	// 注册滑动拼图验证码服务
	factory.RegisterService(Constant.BlockPuzzleCaptcha, CaptchaService.NewBlockPuzzleCaptchaService(factory))

	return factory
}

// GenerateVerificationCode 生成6位随机数字验证码
// @return code 返回6位随机数字验证码
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())

	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10)) // 生成随机数字字符
	}

	return code
}
