package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	GormAdapter "github.com/casbin/gorm-adapter/v3"
	UtilConstant "github.com/herman-hang/herman/app/constants/util"
	"gorm.io/gorm"
)

// InitEnforcer 初始化Casbin模型
// @param string policy 策略
// @param *gorm.DB db 数据库对象
// @return cachedEnforcer err 返回Casbin对象和一个错误信息
func InitEnforcer(policy string, db *gorm.DB) (cachedEnforcer *casbin.CachedEnforcer, err error) {
	// gorm适配器
	adapter, err := GormAdapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	m, err := model.NewModelFromString(policy)
	if err != nil {
		return nil, err
	}
	cachedEnforcer, err = casbin.NewCachedEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	// 设置过期时间为1小时
	cachedEnforcer.SetExpireTime(UtilConstant.ExpireTime)
	_ = cachedEnforcer.LoadPolicy()

	return cachedEnforcer, nil
}

// GetAdminPolicy 管理员策略
// @return string 返回一个策略
func GetAdminPolicy() string {
	return `
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
}
