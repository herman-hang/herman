package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app/common"
	"github.com/herman/app/constants"
	"github.com/herman/app/utils"
	"github.com/mitchellh/mapstructure"
)

// BaseRepository 公共仓储层
type BaseRepository struct {
	Model interface{}
}

// PageInfo 分页结构体
type PageInfo struct {
	Page     int64  `json:"page"`     // 页码
	PageSize int64  `json:"pageSize"` // 每页大小
	Keyword  string `json:"keyword"`  // 关键字
}

// Add 新增
// @param map[string]interface{} data 待添加数据
// @return toMap err 查询数据，错误信息
func (base *BaseRepository) Add(data map[string]interface{}) (toMap map[string]interface{}, err error) {
	data["id"] = constants.InitId
	if err := mapstructure.Decode(data, base.Model); err != nil {
		return nil, err
	}
	if err := common.Db.Create(base.Model).Error; err != nil {
		return nil, err
	}
	// 模型拷贝
	structData := base.Model
	toMap, err = utils.ToMap(structData, "json")
	if err != nil {
		return nil, err
	}
	return toMap, nil
}

// Find 根据ID获取详情
// @param []uint id 主键ID
// @param []string fields 查询指定字段
// @return data err 详情数据，错误信息
func (base *BaseRepository) Find(ids []uint, fields []string) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	if err := common.Db.Model(&base.Model).Select(fields).First(data, ids).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Update 更新
// @param []int ids 查询条件
// @param map[string]interface{} attributes 待更新数据
// @return error 错误信息
func (base *BaseRepository) Update(ids []int, attributes map[string]interface{}) error {
	if err := common.Db.Model(&base.Model).Where("id IN (?)", ids).Updates(attributes).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除
// @param []int ids 主键ID
// @return error 错误信息
func (base *BaseRepository) Delete(ids []int) error {
	if err := common.Db.Delete(&base.Model, ids).Error; err != nil {
		return err
	}
	return nil
}

// IsExist 查询数据是否存在
// @param uint id 条件ID
// @return bool error 返回一个错误信息
func (base *BaseRepository) IsExist(id uint) (bool bool, err error) {
	result := make(map[string]interface{})
	err = common.Db.Model(&base.Model).Find(&result, id).Error
	if len(result) != constants.LengthByZero {
		return true, nil
	}
	return false, err
}

// GetList 获取列表数据
// @param string query 查询条件
// @param []string field 查询指定字段
// @param string order 排序条件
// @return list total pageNum err 返回列表，总条数，总页码数，错误信息
func (base *BaseRepository) GetList(query string, field []string, order string) (data []map[string]interface{}, err error) {
	var (
		ctx     *gin.Context
		page    *PageInfo
		total   int64
		pageNum int64
	)
	// 分页结构体绑定
	if err := ctx.ShouldBindQuery(&page); err != nil {
		return nil, err
	}
	// 总条数
	common.Db.Model(&base.Model).Count(&total)
	// 计算总页数
	pageNum = total / page.PageSize
	if total%page.PageSize != 0 {
		pageNum++
	}
	// 示例 query = fmt.Sprintf(" dns like '%%%s' ", createDbnameInfo.DNS)
	err = common.Db.Model(&base.Model).
		Select(field).
		Where(query).
		Order(order).
		Limit(int(page.PageSize)).
		Offset(int((page.Page - 1) * page.PageSize)).
		Find(&data).Error
	// 向切片追加数据
	data = append(data, map[string]interface{}{
		"total":    total,         // 总条数
		"pageNum":  pageNum,       // 总页数
		"pageSize": page.PageSize, // 每页大小
		"page":     page.Page,     // 当前页码
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetAllData 获取全部数据
// @param []string field 查询指定字段
// @return list err 返回列表，错误信息
func (base *BaseRepository) GetAllData(field []string) (data []map[string]interface{}, err error) {
	if len(field) != 0 {
		if err := common.Db.Model(&base.Model).Select(field).Find(&data).Error; err != nil {
			return nil, err
		}
	}

	if err := common.Db.Model(&base.Model).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
