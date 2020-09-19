package repository

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/utils/lib/databases"
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

// 创建
func (r *User) Creat(user *model.User) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Create(r.dB, user)
	return
}

// Save 会保存所有的字段，即使字段是零值
func (r *User) Save(user *model.User) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Save(r.dB, user)
	return
}

// 修改
func (r *User) Updates(user *model.User, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Updates(r.dB, r.user, user, query, args...)
	return
}

// 删除
func (r *User) Delete(user *model.User, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Delete(r.dB, user, query, args...)
	return
}

func (r *User) First(where interface{}, args ...interface{}) (user *model.User, err error) {
	err = databases.First(r.dB, user, where, args ...)
	return user, err
}

// 根据用户id获取用户数据
func (r *User) FirstById(userId uint64) (user *model.User, err error) {
	err = databases.FirstById(r.dB, user, userId)
	return user, err
}

// 根据用户id获取用户数据
func (r *User) FirstByIdWhere(userId uint64, wheres ...databases.Condition) (user *model.User, err error) {
	err = databases.FirstByIdWhere(r.dB, user, userId, wheres...)
	return user, err
}

// 提交查询
func (r *User) FirstWhere(query interface{}, wheres ...databases.Condition) (user *model.User, err error) {
	err = databases.FirstWhere(r.dB, user, query, wheres...)
	return user, err
}

// 查询用户列表
func (r *User) Find(query interface{}, args ...interface{}) (users []*model.User, err error) {
	err = databases.Find(r.dB, users, query, args...)
	return users, err
}

// 根据用户ids获取用户数据
func (r *User) FindByIds(userIds []uint64) ([]*model.User, error) {
	var (
		err   error
		users = make([]*model.User, 0, len(userIds))
	)
	err = databases.FindByIds(r.dB, users, userIds)
	return users, err
}

// 根据用户ids获取用户数据
func (r *User) FindByIdsWhere(userIds []uint64, wheres ...databases.Condition) ([]*model.User, error) {
	var (
		err   error
		users = make([]*model.User, 0, len(userIds))
	)
	err = databases.FindByIdsWhere(r.dB, users, userIds, wheres...)
	return users, err
}

// 根据用户ids获取用户数据
func (r *User) FindWhere(query interface{}, wheres ...databases.Condition) (users []*model.User, err error) {
	err = databases.FindWhere(r.dB, users, query, wheres...)
	return users, err
}

// 判断邮件是否意见存在 存在
func (r *User) EmailExist(email string) (err error) {
	var (
		user  = new(model.User)
		query = model.User{
			Email: email,
		}
		field = databases.Select("id")
	)
	err = databases.FirstWhere(r.dB, user, query, field)
	return
}
