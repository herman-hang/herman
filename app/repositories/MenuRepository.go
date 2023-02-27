package repositories

import "github.com/herman-hang/herman/app/models"

var Menu = MenuRepository{BaseRepository{Model: new(models.Menu)}}

// MenuRepository 菜单仓储层
type MenuRepository struct {
	BaseRepository
}
