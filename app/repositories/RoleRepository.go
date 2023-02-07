package repositories

import "github.com/herman/app/models"

var Role = &RoleRepository{BaseRepository{Model: new(models.Role)}}

// RoleRepository 角色仓储层
type RoleRepository struct {
	BaseRepository
}
