package repositories

import (
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/models"
)

// DictionaryRepository 数据字典表仓储层
type DictionaryRepository struct {
	BaseRepository
}

// Dictionary 数据字典表仓储层
// @return DictionaryRepository 返回数据字典表仓储层
func Dictionary() *DictionaryRepository {
	return &DictionaryRepository{BaseRepository{Model: new(models.Dictionary)}}
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
