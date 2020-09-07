package service

import (
	"fmt"
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/logger"
)

type User struct {
	user *repository.User
	log  *logger.Logger
}

//
func NewUser(user *repository.User, log *logger.Logger) *User {
	return &User{user: user, log: log}
}

// 生成用户
func (s *User) Create(user *model.User) (int64, error) {
	rowsAffected, err := s.user.Creat(user)
	return rowsAffected, err
}

func (s *User) FirstById(userId uint64) {
	var (
		err  error
		user *model.User
	)
	user, err = s.user.FirstById(userId)
	if err != nil {
		fmt.Println(user)
	}
}
