package models

import (
	"gorm.io/gorm"
	"time"
)

// Admin 管理员结构体
type Admin struct {
	Id           uint           `json:"id" gorm:"column:id;type:uint(11);primary_key;not null;comment:管理员ID"`
	User         string         `json:"user" gorm:"column:user;type:varchar(20);unique;not null;comment:管理员用户名"`
	Password     string         `json:"password" gorm:"column:password;type:char(60);not null;comment:管理员密码"`
	Photo        string         `json:"photo" gorm:"column:photo;type:varchar(255);comment:管理员头像"`
	Name         string         `json:"name" gorm:"column:name;type:varchar(10);comment:真实姓名"`
	Card         string         `json:"card" gorm:"column:card;type:char(20);comment:身份证号码"`
	Sex          uint8          `json:"sex" gorm:"column:sex;type:tinyint(4);default:3;not null;comment:性别(1为女,2为男，3为保密)"`
	Age          uint8          `json:"age" gorm:"column:age;type:tinyint(4);default:0;not null;comment:年龄"`
	Region       string         `json:"region" gorm:"column:region;type:varchar(255);comment:地区"`
	Phone        string         `json:"phone" gorm:"column:phone;type:varchar(16);comment:手机号码"`
	Email        string         `json:"email" gorm:"column:email;type:varchar(32);comment:邮箱"`
	Introduction string         `json:"introduction" gorm:"column:introduction;type:text;comment:简介"`
	State        uint8          `json:"state" gorm:"column:state;type:tinyint(4);default:2;not null;comment:状态(1已停用,2已启用)"`
	Role         uint8          `json:"role" gorm:"column:role;type:varchar(4);not null;comment:角色英文KEY"`
	Sort         uint           `json:"sort" gorm:"column:sort;type:uint(1);default:0;not null;comment:排序"`
	LoginOutIp   string         `json:"loginOutIp" gorm:"column:login_out_ip;type:varchar(32);comment:上一次登录IP地址"`
	LoginTotal   uint           `json:"loginTotal" gorm:"column:login_total;type:uint(11);default:0;not null;comment:登录总数"`
	LoginOutAt   string         `json:"loginOutAt" gorm:"column:login_out_at;type:time;default:1970-01-01 00:00:00;comment:上一次登录时间"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:created_at;type:time;not null;comment:创建时间"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updated_at;type:time;not null;comment:更新时间"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:time;index;comment:删除时间"`
}

// TableName 设置用户表名
func (Admin) TableName() string {
	return "admin"
}
