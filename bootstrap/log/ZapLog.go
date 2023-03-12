package log

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/bootstrap/core"
	"github.com/herman-hang/herman/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// InitZapLogs 初始化日志配置
// @param *settings.Log config 日志配置信息
// @param string mode 当前应用运行模式
// @return err error 返回错误信息
func InitZapLogs(config *config.Log, mode string) (err error) {
	var (
		level = new(zapcore.Level)
		core  zapcore.Core
	)
	// writers
	writersSyncers := GetLoggerWriter(config)
	// 将日志消息编码为指定的格式
	encoders := GetEncoder()
	if err = level.UnmarshalText([]byte(config.Level)); err != nil {
		return err
	}
	if mode == "debug" {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		// 将多个 zapcore.Core 对象合并成一个
		core = zapcore.NewTee(
			zapcore.NewCore(encoders, writersSyncers, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoders, writersSyncers, level)
	}

	logger := zap.New(core, zap.AddCaller())
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logger)
	return nil

}

// GetLoggerWriter return writerSyncer
// @param *settings.Log config 日志配置信息
// @return zapcore.WriteSyncer 返回一个日志记录器
func GetLoggerWriter(config *config.Log) zapcore.WriteSyncer {
	lumberLoggers := &lumberjack.Logger{
		Filename:   core.RootPath + "runtime/logs/" + config.FileName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}
	// zapcore.Lock() 确保多个 goroutine 同时写入日志时不会出现并发问题
	return zapcore.Lock(zapcore.AddSync(lumberLoggers))
}

// GetEncoder 将日志消息编码为指定的格式
// @return zapcore.Encoder 返回Encoder
func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.FunctionKey = "function"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// GinLogger 接收gin框架默认的日志
// @return gin.HandlerFunc 返回中间件上下文
func GinLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		path := ctx.Request.URL.Path
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
