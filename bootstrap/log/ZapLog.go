package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// InitZapLogs 初始化日志配置
// @param *settings.LogConfig config 日志配置信息
// @param string mode 当前应用运行模式
// @return err error 返回错误信息
func InitZapLogs(config *config.LogConfig, mode string) (err error) {
	// writers
	writersSyncers := GetLoggerWriter(config)
	// encoder
	encoders := GetEncoder()
	// 定义接受的l
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.Level))
	if err != nil {
		return err
	}
	var core zapcore.Core
	if mode == "test" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoders, writersSyncers, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoders, writersSyncers, l)
	}

	lg := zap.New(core, zap.AddCaller())
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return

}

// GetLoggerWriter return writerSyncer
// @param *settings.LogConfig config 日志配置信息
// @return zapcore.WriteSyncer 返回一个日志记录器
func GetLoggerWriter(config *config.LogConfig) zapcore.WriteSyncer {
	lumberLoggers := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s%s", "runtime/logs/", config.FileName),
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}
	return zapcore.AddSync(lumberLoggers)
}

// GetEncoder return encoders
// @return zapcore.Encoder 返回Encoder
func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// GinLogger 接收gin框架默认的日志
// @return gin.HandlerFunc 返回中间件上下文
func GinLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		ctx.Next()

		zap.L().Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", time.Since(time.Now())),
		)
	}
}
