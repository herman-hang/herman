package core

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	GormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	UtilConstant "github.com/herman-hang/herman/application/constants/common/util"
	"github.com/herman-hang/herman/kernel/app"
	"github.com/herman-hang/herman/kernel/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// 初始化项目目录
var (
	RootPath     string
	AppPath      string
	ConfigPath   string
	StoragePath  string
	RoutePath    string
	ResourcePath string
)

// init 初始化项目目录
func init() {
	// 获取当前文件所在目录
	dir := absPathing("./")
	// 向上遍历文件目录，直到找到项目根目录
	for {
		if _, err := os.Stat(filepath.Join(dir, "main.go")); err == nil || !os.IsNotExist(err) {
			RootPath = dir
			AppPath = filepath.Join(dir, "app")
			ConfigPath = filepath.Join(dir, "config")
			StoragePath = filepath.Join(dir, "storages")
			RoutePath = filepath.Join(dir, "routes")
			ResourcePath = filepath.Join(dir, "resources")
			break
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			fmt.Println("Can not find project root path")
			os.Exit(1)
		}
		dir = parentDir
	}
}

// absPathing 获取绝对路径
// @param inPath string 当前路径
// @return string 绝对路径
func absPathing(inPath string) string {
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = userHomeDir() + inPath[5:]
	}

	inPath = os.ExpandEnv(inPath)

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}
	return ""
}

// userHomeDir 返回当前用户的主目录
// @return string 主目录
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Db 数据库对象实例
// @return *gorm.DB
func Db() *gorm.DB {
	// 连接Mysql
	db, err := mysql.InitGormDatabase(app.Config.Mysql)
	if err != nil {
		zap.S().Fatal(color.RedString(err.Error()))
	}
	return db
}

// Redis Redis实例对象
// @return *redis.Client
func Redis() *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			app.Config.Redis.Host,
			app.Config.Redis.Port,
		),
		Username: app.Config.Redis.UserName,
		Password: app.Config.Redis.Password,
		DB:       app.Config.Redis.Db,
		PoolSize: app.Config.Redis.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		zap.S().Fatal(color.RedString(err.Error()))
	}

	return rdb
}

// Casbin Casbin对象实例化
// @return *casbin.CachedEnforcer
func Casbin() *casbin.CachedEnforcer {
	db := Db()
	// gorm适配器
	adapter, err := GormAdapter.NewAdapterByDB(db)
	if err != nil {
		zap.S().Fatal(color.RedString(err.Error()))
	}
	policy := `
			[request_definition]
			r = sub, obj, act
			
			[policy_definition]
			p = sub, obj, act
			
			[role_definition]
			g = _, _
			
			[policy_effect]
			e = some(where (p.eft == allow))
			
			[matchers]
			m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`
	m, err := model.NewModelFromString(policy)
	if err != nil {
		zap.S().Fatal(color.RedString(err.Error()))
	}
	cachedEnforcer, err := casbin.NewCachedEnforcer(m, adapter)
	if err != nil {
		zap.S().Fatal(color.RedString(err.Error()))
	}
	// 设置过期时间为1小时
	cachedEnforcer.SetExpireTime(UtilConstant.ExpireTime)
	_ = cachedEnforcer.LoadPolicy()

	return cachedEnforcer
}
