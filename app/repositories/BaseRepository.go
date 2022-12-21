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

type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
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

// Add 批量新增
// @param model 模型结构体
// @return error 返回一个错误信息
func (base *BaseRepository) Add(model []interface{}) error {
	err := common.Db.Create(&model).Error
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
	err := common.Db.Select(fields).First(&base.Model, id).Error
	if err != nil {
		return &base.Model, err
	}
	return &base.Model, nil
}

// Update 批量更新
// @param condition 查询条件
// @param attributes 待更新数据
// @return error 返回一个自定义错误信息
func (base *BaseRepository) Update(condition []int, attributes map[string]interface{}) error {
	err := common.Db.Model(&base.Model).Where("id IN (?)", condition).Updates(attributes).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 批量删除
// @param ids 主键ID
// @return error 返回一个错误信息
func (base *BaseRepository) Delete(ids []int) error {
	err := common.Db.Delete(&base.Model, ids).Error
	if err != nil {
		return err
	}
	return nil
}

// GetList 获取列表数据
// @param query 查询条件
// @param field 查询指定字段
// @param order 排序条件
// @return list total pageNum err 返回列表，总条数，总页码数，错误信息
func (base *BaseRepository) GetList(query string, field []string, order string) (list interface{}, total int, pageNum int, err error) {
	var (
		page  *PageInfo
		model []BaseRepository
	)
	// 查询总页数
	common.Db.Model(&base.Model).Count(&total)
	pageNum = total / page.PageSize
	if total%page.PageSize != 0 {
		pageNum++
	}
	// 示例 query = fmt.Sprintf(" dns like '%%%s' ", createDbnameInfo.DNS)
	err = common.Db.Select(field).Where(query).Order(order).Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Find(&model).Error
	if err != nil {
		return model, total, pageNum, err
	}
	return &model, total, pageNum, nil
}

// GetAllData 获取全部数据
// @param field 查询指定字段
// @return list err 返回列表，错误信息
func (base *BaseRepository) GetAllData(field []string) (list interface{}, err error) {
	var model []BaseRepository
	if len(field) != 0 {
		err = common.Db.Select(field).Find(&model).Error
	}
	err = common.Db.Find(&model).Error
	if err != nil {
		return &model, err
	}
	return &model, nil
}
