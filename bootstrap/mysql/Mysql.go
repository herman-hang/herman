package mysql

import (
	"fmt"
	"github.com/fp/fp-gin-framework/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// InitGormDatabase 初始化gorm数据库
// @param *settings.MysqlConfig config Mysql配置信息
// @return db err 返回一个DB对象和错误信息
func InitGormDatabase(config *config.MysqlConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	// 设置数据库池最大连接数
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	// 设置连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于config.MaxIdleConn，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, err
}
