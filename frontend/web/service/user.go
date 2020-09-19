package service

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/ZRothschild/goIris/config/logger"
	"github.com/ZRothschild/goIris/frontend/web/param/frontendReq"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sync"
)

var (
	onceUser = new(sync.Once)
)

type User struct {
	user *repository.User
	log  *logger.Logger
}

// 单利user service
func NewUser(user *repository.User, log *logger.Logger) (userSrv *User) {
	onceUser.Do(func() {
		userSrv = &User{user: user, log: log}
	})
	return
}

// 用户注册
func (s *User) Register(req *frontendReq.UserRegister) (rowsAffected int64, err error) {
	var (
		user model.User
	)

	if err = copier.Copy(&user, req); err != nil {
		return
	}

	// 查找用户是否存在
	if err = s.user.EmailExist(user.Email); err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if err != gorm.ErrRecordNotFound {
		return 0, conf.ErrRecordExist
	}

	rowsAffected, err = s.user.Creat(&user)

	return rowsAffected, err
}

// 生成用户
func (s *User) Create(user *model.User) (rowsAffected int64, err error) {
	rowsAffected, err = s.user.Creat(user)
	return rowsAffected, err
}

// 通过用户id 查找用户
func (s *User) FirstById(userId uint64) (user *model.User, err error) {
	user, err = s.user.FirstById(userId)
	return
}
