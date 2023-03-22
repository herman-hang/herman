package system

import (
	SystemConstant "github.com/herman-hang/herman/app/constants/system"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/bootstrap/core"
)

// Find 获取系统设置信息
// @return map[string]interface{} 返回系统设置信息
func Find() map[string]interface{} {
	fields := []string{
		"name",
		"title",
		"description",
		"keywords",
		"logo_file_id",
		"ico_file_id",
		"record",
		"copyright",
		"is_website",
		"email",
		"telephone",
		"address",
	}
	info, err := repositories.System().Find(map[string]interface{}{"id": SystemConstant.Id}, fields)
	if err != nil {
		core.Debug(err)
		panic(SystemConstant.FindFail)
	}
	return info
}

// Modify 修改系统设置信息
// @param map[string]interface{} data 待修改信息
// @return void
func Modify(data map[string]interface{}) {
	if err := repositories.System().Update([]uint{SystemConstant.Id}, data); err != nil {
		panic(SystemConstant.ModifyFail)
	}
}
