package repository

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/config/db"
	"gorm.io/gorm"
)

// 用户
type User struct {
	user *model.User
	dB   *gorm.DB
}

// 用户示例
func NewUser(user *model.User, dB *gorm.DB) *User {
	return &User{
		user: user,
		dB:   dB,
	}
}

// 根据用户id获取用户数据
func (r *User) FirstById(userId uint64) (*model.User, error) {
	var (
		err  error
		user = new(model.User)
	)
	err = db.FirstById(r.dB, user, userId)
	return user, err
}
