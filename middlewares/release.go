package middlewares

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	CasbinServer "github.com/herman-hang/herman/kernel/casbin"
	"github.com/herman-hang/herman/kernel/core"
	"github.com/herman-hang/herman/kernel/mysql"
	RedisServer "github.com/herman-hang/herman/kernel/redis"
	"github.com/herman-hang/herman/servers/settings"
	"go.uber.org/zap"
)

// ServerHandler 服务管理中间件
// @return gin.HandlerFunc
func ServerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Reload()
		ctx.Next()
		Close()
	}
}

// Reload 加载服务函数
// @return void
func Reload() {
	// 连接Mysql
	db, err := mysql.InitGormDatabase(settings.Config.Mysql)
	if err != nil {
		zap.S().Fatal(color.RedString(fmt.Sprintf("Mysql connection failure:%v", err)))
	}

	// 连接Redis
	rdb, err := RedisServer.InitRedisConfig(settings.Config.Redis)
	if err != nil {
		zap.S().Fatal(color.RedString(fmt.Sprintf("Redis connection failed:%v", err)))
	}

	// 初始化Casbin
	cachedEnforcer, err := CasbinServer.InitEnforcer(CasbinServer.GetAdminPolicy(), db)
	if err != nil {
		zap.S().Fatal(color.RedString(fmt.Sprintf("The casbin initialization failed:%v", err)))
	}

	core.Db, core.Redis, core.Casbin = db, rdb, cachedEnforcer
}

// Close 释放服务函数
// @return void
func Close() {
	// 关闭redis
	_ = core.Redis.Close()
	// 关闭DB
	db, _ := core.Db.DB()
	if db != nil {
		_ = db.Close()
	}
	// 设置全局变量为nil，等待GC进行回收
	core.Db, core.Redis, core.Casbin = nil, nil, nil
}
