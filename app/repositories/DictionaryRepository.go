package repositories

import (
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/models"
)

// Dictionary 实例化结构体并重写BaseRepository
var Dictionary = DictionaryRepository{BaseRepository{Model: new(models.Dictionary)}}

// DictionaryRepository 数据字典表仓储层
type DictionaryRepository struct {
	BaseRepository
}

// DeleteByDictionaryId 根据数据字典ID删除明细值
// @param []uint id 数据字典ID
// @return error 返回一个错误信息
func (base DictionaryRepository) DeleteByDictionaryId(id []uint) error {
	err := common.Db.Where("dictionary_id IN ?", id).Delete(&base.Model).Error
	if err != nil {
		return err
	}
	return nil
}
