package captcha

import (
	"fmt"
	captchaConfig "github.com/TestsLing/aj-captcha-go/config"
	constant "github.com/TestsLing/aj-captcha-go/const"
	captchaService "github.com/TestsLing/aj-captcha-go/service"
	"github.com/fp/fp-gin-framework/config"
)

// InitCaptcha 初始化滑块验证码
func InitCaptcha(config *config.AppConfig) (factory *captchaService.CaptchaServiceFactory) { // 行为校验配置模块（具体参数可从业务系统配置文件自定义）
	// 行为校验初始化
	factory = captchaService.NewCaptchaServiceFactory(
		captchaConfig.BuildConfig(config.CaptchaConfig.CacheType,
			config.CaptchaConfig.ResourcePath,
			&captchaConfig.WatermarkConfig{
				Text: config.CaptchaConfig.Text,
			},
			nil, nil, config.CaptchaConfig.CacheExpireSec))
	// 注册内存缓存
	factory.RegisterCache(constant.MemCacheKey, captchaService.NewMemCacheService(20))
	// 注册自定义配置redis数据库
	factory.RegisterCache(constant.RedisCacheKey, captchaService.NewConfigRedisCacheService([]string{fmt.Sprintf("%s:%d",
		config.RedisConfig.Host,
		config.RedisConfig.Port,
	)},
		config.RedisConfig.UserName,
		config.RedisConfig.Password,
		false,
		config.RedisConfig.Db,
	))
	// 注册文字点选验证码服务
	factory.RegisterService(constant.ClickWordCaptcha, captchaService.NewClickWordCaptchaService(factory))
	// 注册滑动拼图验证码服务
	factory.RegisterService(constant.BlockPuzzleCaptcha, captchaService.NewBlockPuzzleCaptchaService(factory))
	return factory
}
