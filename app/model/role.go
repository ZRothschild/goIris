package model

import (
	"github.com/ZRothschild/goIris/app/model/base"
)

// 角色表
type Role struct {
	base.Base
	Name     string `gorm:"not null;default:'';type:varchar(60);column:name;comment:角色名称" json:"name"`
	Status   uint8  `gorm:"not null;default:1;column:status;comment:1 正常 2 非正常" json:"status"`
	RoleType uint8  `gorm:"not null;default:1;column:role_type;comment:1 普通  2 组" json:"roleType"`
}

// 角色
func NewRole() *Role {
	return &Role{}
}

// 表名
func (m *Role) TableName() string {
	return "roles"
}
