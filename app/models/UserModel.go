package models

import (
	"time"
)

type Users struct {
	Id           uint       `json:"id" gorm:"primary_key" gorm:"comment:主键ID"`
	User         string     `json:"user" gorm:"comment:用户名"`
	Password     string     `json:"password" gorm:"comment:用户密码"`
	Photo        string     `json:"photo" gorm:"comment:用户头像"`
	Nickname     string     `json:"nickname" gorm:"comment:昵称"`
	Name         string     `json:"name" gorm:"comment:真实姓名"`
	Card         string     `json:"card" gorm:"comment:身份证号码"`
	Sex          string     `json:"sex" gorm:"comment:性别(0为女,1为男，2为保密)"`
	Age          int        `json:"age" gorm:"comment:年龄"`
	Region       string     `json:"region" gorm:"comment:地区"`
	Phone        string     `json:"phone" gorm:"comment:手机号码"`
	Email        string     `json:"email" gorm:"comment:邮箱"`
	Introduction string     `json:"introduction" gorm:"comment:简介"`
	Status       string     `json:"status" gorm:"comment:状态(0已停用,1已启用)"`
	SignOutIp    string     `json:"sign_out_ip" gorm:"comment:最后登录IP地址"`
	SignTotal    string     `json:"sign_total" gorm:"comment:登录总数"`
	SignOutAt    string     `json:"sign_out_at" gorm:"comment:最后登录时间"`
	CreatedAt    time.Time  `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"comment:更新时间"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index" gorm:"comment:删除时间"`
}
