package repositories

import (
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// DictionaryRepository 数据字典表仓储层
type DictionaryRepository struct {
	BaseRepository
}

// Dictionary 数据字典表仓储层
// @param *gorm.DB tx 事务
// @return DictionaryRepository 返回数据字典表仓储层
func Dictionary(tx ...*gorm.DB) *DictionaryRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &DictionaryRepository{BaseRepository{Model: new(models.Dictionary), Db: tx[0]}}
	}
	return &DictionaryRepository{BaseRepository{Model: new(models.Dictionary), Db: core.Db}}
}

// DeleteByDictionaryId 根据数据字典ID删除明细值
// @param []uint id 数据字典ID
// @return error 返回一个错误信息
func (base DictionaryRepository) DeleteByDictionaryId(id []uint) error {
	err := base.Db.Where("dictionary_id IN ?", id).Delete(&base.Model).Error
	if err != nil {
		return err
	}
	return nil
}
