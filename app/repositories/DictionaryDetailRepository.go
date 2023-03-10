package repositories

import (
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/models"
)

// DictionaryDetailRepository 数据字典表仓储层
type DictionaryDetailRepository struct {
	BaseRepository
}

// DictionaryDetail 数据字典表仓储层
// @return DictionaryDetailRepository 返回数据字典表仓储层
func DictionaryDetail() *DictionaryDetailRepository {
	return &DictionaryDetailRepository{BaseRepository{Model: new(models.DictionaryDetail)}}
}

// FindByCode 根据数据字典KEY返回明细值
// @param map[string]interface{} condition 查询条件
// @param []string fields 查询指定字段
// @return data err 切片数据，错误信息
func (base DictionaryDetailRepository) FindByCode(condition map[string]interface{}, fields ...[]string) (data []map[string]interface{}, err error) {
	var list []map[string]interface{}
	if err := common.Db.Model(base.Model).Where(condition).Select(fields[0]).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
