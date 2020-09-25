package model

import (
	"github.com/ZRothschild/goIris/app/model/base"
)

// 权限表
type Permission struct {
	base.Base
	Name           string `gorm:"not null;default:'';type:varchar(60);column:name;comment:权限名称" json:"name"`
	Url            string `gorm:"not null;default:'';type:varchar(80);column:url;comment:路径" json:"url"`
	Method         string `gorm:"not null;default:'';type:varchar(20);column:method;comment:路径" json:"method"`
	PermissionType uint8  `gorm:"not null;default:1;column:permission_type;comment:1 普通 2 组" json:"permissionType"`
}

// 权限
func NewPermission() *Permission {
	return &Permission{}
}

// 表名
func (m *Permission) TableName() string {
	return "permissions"
}
