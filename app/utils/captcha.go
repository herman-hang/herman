package utils

import (
	"fmt"
	CaptchaConfig "github.com/TestsLing/aj-captcha-go/config"
	Constant "github.com/TestsLing/aj-captcha-go/const"
	CaptchaService "github.com/TestsLing/aj-captcha-go/service"
	CaptchaConstant "github.com/herman-hang/herman/app/constants/captcha"
	"github.com/herman-hang/herman/servers/settings"
)

// Factory 初始化滑块验证码
// @return factory 返回一个验证码工厂
func Factory() (factory *CaptchaService.CaptchaServiceFactory) { // 行为校验配置模块（具体参数可从业务系统配置文件自定义）
	// 行为校验初始化
	factory = CaptchaService.NewCaptchaServiceFactory(
		CaptchaConfig.BuildConfig(settings.Config.Captcha.CacheType,
			settings.Config.Captcha.ResourcePath,
			&CaptchaConfig.WatermarkConfig{
				Text: settings.Config.Captcha.Text,
			},
			nil, nil, settings.Config.Captcha.CacheExpireSec))
	// 注册内存缓存
	factory.RegisterCache(Constant.MemCacheKey, CaptchaService.NewMemCacheService(CaptchaConstant.CacheMaxNumber))
	// 注册自定义配置redis数据库
	factory.RegisterCache(Constant.RedisCacheKey, CaptchaService.NewConfigRedisCacheService([]string{fmt.Sprintf("%s:%d",
		settings.Config.Redis.Host,
		settings.Config.Redis.Port,
	)},
		settings.Config.Redis.UserName,
		settings.Config.Redis.Password,
		false,
		settings.Config.Redis.Db,
	))
	// 注册文字点选验证码服务
	factory.RegisterService(Constant.ClickWordCaptcha, CaptchaService.NewClickWordCaptchaService(factory))
	// 注册滑动拼图验证码服务
	factory.RegisterService(Constant.BlockPuzzleCaptcha, CaptchaService.NewBlockPuzzleCaptchaService(factory))

	return factory
}
