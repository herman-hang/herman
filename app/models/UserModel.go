package models

import (
	"gorm.io/gorm"
	"time"
)

// Users 用户表结构体
type Users struct {
	Id           uint           `json:"id" gorm:"column:id;primary_key;comment:用户ID"`
	User         string         `json:"user" gorm:"column:user;comment:用户名"`
	Password     string         `json:"password" gorm:"column:password;comment:用户密码"`
	Photo        string         `json:"photo" gorm:"column:photo;comment:用户头像"`
	Nickname     string         `json:"nickname" gorm:"column:nickname;comment:昵称"`
	Name         string         `json:"name" gorm:"column:name;comment:真实姓名"`
	Card         string         `json:"card" gorm:"column:card;comment:身份证号码"`
	Sex          uint8          `json:"sex" gorm:"column:sex;default:3;comment:性别(1为女,2为男，3为保密)"`
	Age          uint8          `json:"age" gorm:"column:age;default:0;comment:年龄"`
	Region       string         `json:"region" gorm:"column:region;comment:地区"`
	Phone        string         `json:"phone" gorm:"column:phone;comment:手机号码"`
	Email        string         `json:"email" gorm:"column:email;comment:邮箱"`
	Introduction string         `json:"introduction" gorm:"column:introduction;comment:简介"`
	State        uint8          `json:"state" gorm:"column:state;default:2;comment:状态(1已停用,2已启用)"`
	Sort         uint           `json:"sort" gorm:"column:sort;default:0;comment:排序"`
	LoginOutIp   string         `json:"loginOutIp" gorm:"column:login_out_ip;comment:上一次登录IP地址"`
	LoginTotal   string         `json:"loginTotal" gorm:"column:login_total;default:0;comment:登录总数"`
	LoginOutAt   string         `json:"loginOutAt" gorm:"column:login_out_at;comment:上一次登录时间"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置用户表名
func (Users) TableName() string {
	return "users"
}
