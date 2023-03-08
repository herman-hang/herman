package repositories

import (
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/constants"
	"github.com/herman-hang/herman/app/utils"
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
	Keywords string `json:"keywords"` // 关键字
}

// Insert 新增
// @param map[string]interface{} data 待添加数据
// @return toMap err 查询数据，错误信息
func (base *BaseRepository) Insert(data map[string]interface{}) (toMap map[string]interface{}, err error) {
	// 初始化ID，让ID持续自增
	data["id"] = constants.InitId
	if err := mapstructure.WeakDecode(data, base.Model); err != nil {
		return nil, err
	}
	if err := common.Db.Create(base.Model).Error; err != nil {
		return nil, err
	}
	// 模型拷贝
	tempStruct := base.Model
	toMap, err = utils.ToMap(tempStruct, "json")
	if err != nil {
		return nil, err
	}
	return toMap, nil
}

// Find 根据查询条件获取详情
// @param map[string]interface{} condition 查询条件
// @param []string fields 查询指定字段
// @return data err 详情数据，错误信息
func (base *BaseRepository) Find(condition map[string]interface{}, fields ...[]string) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	if len(fields) > 0 {
		if err := common.Db.Model(&base.Model).Where(condition).Select(fields[0]).Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		if err := common.Db.Model(&base.Model).Where(condition).Find(&data).Error; err != nil {
			return nil, err
		}
	}
	return data, nil
}

// Update 更新
// @param []uint ids 查询条件
// @param map[string]interface{} attributes 待更新数据
// @return error 错误信息
func (base *BaseRepository) Update(ids []uint, data map[string]interface{}) error {
	var attributes = make(map[string]interface{})
	// 驼峰转下划线
	for k, v := range data {
		k := utils.ToSnakeCase(k)
		attributes[k] = v
	}
	if err := common.Db.Model(&base.Model).Where("id IN (?)", ids).Updates(attributes).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除
// @param []uint ids 主键ID
// @return error 错误信息
func (base *BaseRepository) Delete(ids []uint) error {
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
// @param []string fields 查询指定字段
// @param string order 排序条件
// @param map[string]interface{} pageInfo 列表分页和关键词数据
// @return list total pageNum err 返回列表，总条数，总页码数，错误信息
func (base *BaseRepository) GetList(query string, fields []string, order string, pageInfo ...map[string]interface{}) (data map[string]interface{}, err error) {
	var (
		page    PageInfo
		total   int64
		pageNum int64
		list    []map[string]interface{}
	)
	if len(pageInfo) > 0 {
		if err := mapstructure.WeakDecode(pageInfo[0], &page); err != nil {
			panic(constants.MapToStruct)
		}
	}
	// 总条数
	common.Db.Model(&base.Model).Count(&total)
	// 计算总页数
	if page.PageSize != 0 && total%page.PageSize != 0 {
		pageNum = total / page.PageSize
		pageNum++
	}
	// 示例 query = fmt.Sprintf(" dns like '%%%s' ", createDbnameInfo.DNS)
	err = common.Db.Model(&base.Model).
		Select(fields).
		Where(query).
		Order(order).
		Limit(int(page.PageSize)).
		Offset(int((page.Page - 1) * page.PageSize)).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	data = map[string]interface{}{
		"list":     list,          // 数据
		"total":    total,         // 总条数
		"pageNum":  pageNum,       // 总页数
		"pageSize": page.PageSize, // 每页大小
		"page":     page.Page,     // 当前页码
	}
	return data, nil
}

// GetAllData 获取全部数据
// @param []string fields 查询指定字段
// @return list err 返回列表，错误信息
func (base *BaseRepository) GetAllData(fields []string) (data []map[string]interface{}, err error) {
	if len(fields) > 0 {
		if err := common.Db.Model(&base.Model).Select(fields).Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		if err := common.Db.Model(&base.Model).Find(&data).Error; err != nil {
			return nil, err
		}
	}
	return data, nil
}
