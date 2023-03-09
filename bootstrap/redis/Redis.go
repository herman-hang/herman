package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/herman-hang/herman/config"
	"time"
)

// InitRedisConfig 初始化Redis
// @param *settings.Redis cfg Mysql配置信息
// @return rdb err 返回一个redis对象和错误信息
func InitRedisConfig(cfg *config.Redis) (rdb *redis.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Username: cfg.UserName,
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return rdb, err
	}
	return rdb, nil
}
