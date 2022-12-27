package mysql

import (
	"fmt"
	"github.com/fp/fp-gin-framework/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// InitGormDatabase 初始化gorm数据库
// @param *settings.MysqlConfig config Mysql配置信息
// @return *gorm.DB error 返回一个DB对象和错误信息
func InitGormDatabase(config *config.MysqlConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	db, err = gorm.Open("mysql", dsn)
	db.SingularTable(true) // 设置禁用表名复数形式
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Prefix + defaultTableName
	}
	db.DB().SetConnMaxLifetime(time.Minute * 3)
	// 设置数据库池最大连接数
	db.DB().SetMaxOpenConns(config.MaxOpenConn)
	// 设置连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于config.MaxIdleConn，超过的连接会被连接池关闭
	db.DB().SetMaxIdleConns(config.MaxIdleConn)
	return db, err
}
