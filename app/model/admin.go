package model

import (
	"github.com/ZRothschild/goIris/app/model/base"
)

// 管理员表
type Admin struct {
	base.Base
	Nickname string `gorm:"not null;default:'';type:varchar(60);column:nickname;comment:昵称" json:"nickname"`
	Email    string `gorm:"not null;default:'';type:varchar(60);column:email;comment:用户email;index" json:"email"`
	Password string `gorm:"not null;default:'';type:varchar(90);column:password;comment:密码" json:"password"`
	Status   uint8  `gorm:"not null;default:1;column:status;comment:1 正常 2 非正常用户" json:"status"`
}

// 管理员
func NewAdmin() *Admin {
	return &Admin{}
}

// 表名
func (m *Admin) TableName() string {
	return "admins"
}
