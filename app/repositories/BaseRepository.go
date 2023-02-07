package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app/common"
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

// Add 批量新增
// @param []interface{} model 模型结构体
// @return bool 返回一个bool值
func (base *BaseRepository) Add(model []interface{}) bool {
	err := common.Db.Create(&model).Error
	if err != nil {
		return false
	}
	return true
}

// Find 根据ID获取详情
// @param []uint id 主键ID
// @param []string fields 查询指定字段
// @return data, bool 详情数据，bool值
func (base *BaseRepository) Find(ids []uint, fields []string) (data map[string]interface{}, bool bool) {
	err := common.Db.Model(&base.Model).Select(fields).First(data, ids).Error
	if err != nil {
		return nil, false
	}
	return data, true
}

// Update 批量更新
// @param []int ids 查询条件
// @param map[string]interface{} attributes 待更新数据
// @return bool 返回一个bool值
func (base *BaseRepository) Update(ids []int, attributes map[string]interface{}) bool {
	err := common.Db.Model(&base.Model).Where("id IN (?)", ids).Updates(attributes).Error
	if err != nil {
		return false
	}
	return true
}

// Delete 批量删除
// @param []int ids 主键ID
// @return error 返回一个错误信息
func (base *BaseRepository) Delete(ids []int) bool {
	err := common.Db.Delete(&base.Model, ids).Error
	if err != nil {
		return false
	}
	return true
}

// IsExist 查询数据是否存在
// @param uint id 条件ID
// @return bool 返回一个bool值
func (base *BaseRepository) IsExist(id uint) bool {
	result := map[string]interface{}{}
	common.Db.Model(&base.Model).First(&result, id)
	if result != nil {
		return true
	}
	return false
}

// GetList 获取列表数据
// @param string query 查询条件
// @param []string field 查询指定字段
// @param string order 排序条件
// @return list total pageNum bool 返回列表，总条数，总页码数，bool值
func (base *BaseRepository) GetList(query string, field []string, order string) (data []map[string]interface{}, bool bool) {
	var (
		ctx     *gin.Context
		page    *PageInfo
		total   int64
		pageNum int64
	)

	// 分页结构体绑定
	if err := ctx.ShouldBindQuery(&page); err != nil {
		return nil, false
	}
	// 总条数
	common.Db.Model(&base.Model).Count(&total)
	// 计算总页数
	pageNum = total / page.PageSize
	if total%page.PageSize != 0 {
		pageNum++
	}
	// 示例 query = fmt.Sprintf(" dns like '%%%s' ", createDbnameInfo.DNS)
	err := common.Db.Model(&base.Model).
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
		return nil, false
	}
	return data, true
}

// GetAllData 获取全部数据
// @param []string field 查询指定字段
// @return list bool 返回列表，bool值
func (base *BaseRepository) GetAllData(field []string) (data []map[string]interface{}, bool bool) {
	if len(field) != 0 {
		if err := common.Db.Model(&base.Model).Select(field).Find(&data).Error; err != nil {
			return nil, false
		}
	}

	if err := common.Db.Model(&base.Model).Find(&data).Error; err != nil {
		return nil, false
	}
	return data, true
}
