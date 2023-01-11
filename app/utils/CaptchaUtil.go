package utils

import (
	"fmt"
	captchaConfig "github.com/TestsLing/aj-captcha-go/config"
	constant "github.com/TestsLing/aj-captcha-go/const"
	captchaService "github.com/TestsLing/aj-captcha-go/service"
	"github.com/fp/fp-gin-framework/servers/settings"
)

// Factory 初始化滑块验证码
// @return factory 返回一个验证码工厂
func Factory() (factory *captchaService.CaptchaServiceFactory) { // 行为校验配置模块（具体参数可从业务系统配置文件自定义）
	// 行为校验初始化
	factory = captchaService.NewCaptchaServiceFactory(
		captchaConfig.BuildConfig(settings.Config.CaptchaConfig.CacheType,
			settings.Config.CaptchaConfig.ResourcePath,
			&captchaConfig.WatermarkConfig{
				Text: settings.Config.CaptchaConfig.Text,
			},
			nil, nil, settings.Config.CaptchaConfig.CacheExpireSec))
	// 注册内存缓存
	factory.RegisterCache(constant.MemCacheKey, captchaService.NewMemCacheService(20))
	// 注册自定义配置redis数据库
	factory.RegisterCache(constant.RedisCacheKey, captchaService.NewConfigRedisCacheService([]string{fmt.Sprintf("%s:%d",
		settings.Config.RedisConfig.Host,
		settings.Config.RedisConfig.Port,
	)},
		settings.Config.RedisConfig.UserName,
		settings.Config.RedisConfig.Password,
		false,
		settings.Config.RedisConfig.Db,
	))
	// 注册文字点选验证码服务
	factory.RegisterService(constant.ClickWordCaptcha, captchaService.NewClickWordCaptchaService(factory))
	// 注册滑动拼图验证码服务
	factory.RegisterService(constant.BlockPuzzleCaptcha, captchaService.NewBlockPuzzleCaptchaService(factory))

	return factory
}
