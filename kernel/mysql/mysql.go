package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/config"
	"github.com/herman-hang/herman/kernel/app"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// Writer 自定义日志输出
type Writer struct{}

// InitGormDatabase 初始化gorm数据库
// @param *app.Mysql config Mysql配置信息
// @return db err 返回一个DB对象和错误信息
func InitGormDatabase(config *config.Mysql) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	// 创建日志对象
	newLogger := logger.New(
		Writer{}, // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	sqlDB, err := db.DB()
	// 设置数据库池最大连接数
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	// 设置连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于config.MaxIdleConn，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, err
}

// Printf 实现zap的Logger接口
// @param string format
// @param ...interface{} args
// @return void
func (w Writer) Printf(format string, args ...interface{}) {
	if app.Config.Mode == gin.DebugMode {
		zap.L().Info(fmt.Sprintf(format, args...))
	}
}
