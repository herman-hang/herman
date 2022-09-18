package redis

import (
	"fmt"
	"fp-back-user/settings"
	"github.com/go-redis/redis"
)

// InitRedisConfig 初始化Redis
func InitRedisConfig(cfg *settings.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return rdb, err
	}
	return rdb, nil
}
