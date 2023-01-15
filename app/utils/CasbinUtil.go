package utils

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	GormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/fp/fp-gin-framework/app/common"
	UtilConstant "github.com/fp/fp-gin-framework/app/constants/util"
)

func Enforcer(policy string) (cachedEnforcer *casbin.CachedEnforcer) {
	// gorm适配器
	adapter, _ := GormAdapter.NewAdapterByDB(common.Db)

	m, err := model.NewModelFromString(policy)
	if err != nil {
		panic(UtilConstant.StringModelLoadFail)
	}
	cachedEnforcer, err = casbin.NewCachedEnforcer(m, adapter)
	if err != nil {
		panic(UtilConstant.CachedEnforcerFail)
	}
	// 设置过期时间为1小时
	cachedEnforcer.SetExpireTime(60 * 60)
	_ = cachedEnforcer.LoadPolicy()

	return cachedEnforcer
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
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
}
