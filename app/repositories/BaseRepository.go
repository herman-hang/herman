package repositories

import (
	"encoding/json"
	"github.com/fp/fp-gin-framework/app/common"
	"github.com/fp/fp-gin-framework/app/models"
	"github.com/jinzhu/gorm"
)

type BaseRepository struct {
	Model interface{}
}

// UserInfo 根据ID获取用户信息
// @param uint id 用户id
// @return map[string]interface{} 返回当前用户ID的信息
func UserInfo(id uint) map[string]interface{} {
	var users models.Users
	err := common.Db.Where("id = ?", id).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err.Error())
	}

	data, _ := json.Marshal(&users)
	userMap := make(map[string]interface{})
	_ = json.Unmarshal(data, &userMap)

	return userMap
}

// Add 新增操作
// @param model 模型结构体
// @return error 返回一个错误信息
func (base *BaseRepository) Add(model gorm.Model) error {
	err := common.Db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

// Find 根据ID获取详情
// @param id 主键ID
// @param fields 查询指定字段
// @return gorm.Model, error 返回查询的模型和一个错误信息
func (base *BaseRepository) Find(id uint, fields []string) (interface{}, error) {
	err := common.Db.Select(fields).Find(base.Model, id).Error
	if err != nil {
		return base.Model, err
	}
	return base.Model, nil
}
