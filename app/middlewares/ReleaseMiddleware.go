package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app/common"
	CasbinServer "github.com/herman/bootstrap/casbin"
	"github.com/herman/bootstrap/mysql"
	RedisServer "github.com/herman/bootstrap/redis"
	"github.com/herman/servers/settings"
	"go.uber.org/zap"
)

// ServerHandler 服务管理中间件
// @return gin.HandlerFunc
func ServerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Reload(ctx)
		ctx.Next()
		Close()
	}
}

// Reload 加载服务函数
// @return void
func Reload(ctx *gin.Context) {
	// 连接Mysql
	db, err := mysql.InitGormDatabase(settings.Config.MysqlConfig)
	if err != nil {
		zap.S().Errorf("Mysql connection failure:%v", err)
	}

	// 连接Redis
	rdb, err := RedisServer.InitRedisConfig(settings.Config.RedisConfig)
	if err != nil {
		zap.S().Fatalf("Redis connection failed:%v", err)
	}

	// 初始化Casbin
	cachedEnforcer, err := CasbinServer.InitEnforcer(CasbinServer.GetAdminPolicy(), db)
	if err != nil {
		zap.S().Fatalf("The Casbin initialization failed:%v", err)
	}

	common.Db, common.Redis, common.Casbin, common.Context = db, rdb, cachedEnforcer, ctx
}

// Close 释放服务函数
// @return void
func Close() {
	// 释放redis
	_ = common.Redis.Close()
	// 释放DB
	db, _ := common.Db.DB()
	if db != nil {
		_ = db.Close()
	}
	// 释放Casbin
	common.Casbin = nil
}
