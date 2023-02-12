package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app/common"
	"github.com/herman/app/constants"
	"github.com/herman/bootstrap/casbin"
	"github.com/herman/bootstrap/mysql"
	"github.com/herman/servers/settings"
	"gorm.io/gorm"
)

// DbTransactionAfter Db事务结束后事件
// @return void
func DbTransactionAfter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 注册Db事务执行后事件
		_ = common.Db.Callback().Create().After("gorm:commit_or_rollback_transaction").Register("on_commit", func(tx *gorm.DB) {
			Db, _ := mysql.InitGormDatabase(settings.Config.MysqlConfig)
			_, err := casbin.InitEnforcer(casbin.GetAdminPolicy(), Db)
			common.Db = Db
			if err != nil {
				panic(constants.CasbinAdapterByDbFail)
			}
		})
		ctx.Next()
	}
}
