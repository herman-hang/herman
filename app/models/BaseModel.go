package models

import (
	"encoding/json"
	"fp-back-user/app/common"
	"github.com/jinzhu/gorm"
)

// UserInfo 根据ID获取用户信息
// @param uint id 用户id
// @return map[string]interface{} 返回当前用户ID的信息
func UserInfo(id uint) map[string]interface{} {
	var users Users
	err := common.Db.Where("id = ?", id).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err.Error())
	}

	data, _ := json.Marshal(&users)
	userMap := make(map[string]interface{})
	_ = json.Unmarshal(data, &userMap)

	return userMap
}
