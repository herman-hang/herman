package mysql

import (
	"fmt"
	"fp-back-user/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitGormDatabase 初始化gorm数据库
// @param *settings.MysqlConfig config Mysql配置信息
// @return *gorm.DB error 返回一个DB对象和错误信息
func InitGormDatabase(config *settings.MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	db, err := gorm.Open("mysql", dsn)
	db.SingularTable(true) // 设置禁用表名复数形式
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Prefix + defaultTableName
	}
	db.DB().SetMaxIdleConns(config.MaxIdsConn)
	db.DB().SetMaxOpenConns(config.MaxOpenConn)
	return db, err
}
