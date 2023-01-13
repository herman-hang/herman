package utils

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/fp/fp-gin-framework/app/common"
	utilConstant "github.com/fp/fp-gin-framework/app/constants/util"
)

func Enforcer() (cachedEnforcer *casbin.CachedEnforcer) {
	// gorm适配器
	adapter, _ := gormAdapter.NewAdapterByDB(common.Db)
	text := `
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
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(utilConstant.StringModelLoadFail)
	}
	cachedEnforcer, err = casbin.NewCachedEnforcer(m, adapter)
	if err != nil {
		panic(utilConstant.CachedEnforcerFail)
	}
	// 设置过期时间为1小时
	cachedEnforcer.SetExpireTime(60 * 60)
	_ = cachedEnforcer.LoadPolicy()

	return cachedEnforcer
}
