package role

import (
	"github.com/herman/app/common"
	AdminConstant "github.com/herman/app/constants/role"
	"github.com/herman/app/repositories"
	"gorm.io/gorm"
)

func Add(data map[string]interface{}) {
	// 函数结束后归还db对象
	db := common.Db
	defer func() {
		common.Db = db
	}()

	_ = common.Db.Transaction(func(tx *gorm.DB) error {
		common.Db = tx
		// 添加数据
		if !repositories.Role.Add(data) {
			panic(AdminConstant.AddFail)
		}
		if data["pid"].(uint) != AdminConstant.PidNotExist {
			// casbin添加策略
		}
		return nil
	})

}
