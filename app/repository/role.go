package repository

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/utils/lib/databases"
	"gorm.io/gorm"
)

// 用户
type Role struct {
	role *model.Role
	dB   *gorm.DB
}

// 用户示例
func NewRole(role *model.Role, dB *gorm.DB) *Role {
	return &Role{
		role: role,
		dB:   dB,
	}
}

// 创建
func (r *Role) Creat(role *model.Role) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Create(r.dB, role)
	return
}

// Save 会保存所有的字段，即使字段是零值
func (r *Role) Save(role *model.Role) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Save(r.dB, role)
	return
}

// 修改
func (r *Role) Updates(role *model.Role, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Updates(r.dB, r.role, role, query, args...)
	return
}

// 删除
func (r *Role) Delete(role *model.Role, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Delete(r.dB, role, query, args...)
	return
}

func (r *Role) First(where interface{}, args ...interface{}) (role *model.Role, err error) {
	err = databases.First(r.dB, role, where, args...)
	return role, err
}

func (r *Role) FirstById(roleId uint64) (role *model.Role, err error) {
	err = databases.FirstById(r.dB, role, roleId)
	return role, err
}

func (r *Role) FirstByIdWhere(roleId uint64, wheres ...databases.Condition) (role *model.Role, err error) {
	err = databases.FirstByIdWhere(r.dB, role, roleId, wheres...)
	return role, err
}

// 提交查询
func (r *Role) FirstWhere(query interface{}, wheres ...databases.Condition) (role *model.Role, err error) {
	err = databases.FirstWhere(r.dB, role, query, wheres...)
	return role, err
}

func (r *Role) Find(query interface{}, args ...interface{}) (roles []*model.Role, err error) {
	err = databases.Find(r.dB, roles, query, args...)
	return roles, err
}

func (r *Role) FindByIds(roleIds []uint64) ([]*model.Role, error) {
	var (
		err   error
		roles = make([]*model.Role, 0, len(roleIds))
	)
	err = databases.FindByIds(r.dB, roles, roleIds)
	return roles, err
}

func (r *Role) FindByIdsWhere(roleIds []uint64, wheres ...databases.Condition) ([]*model.Role, error) {
	var (
		err   error
		roles = make([]*model.Role, 0, len(roleIds))
	)
	err = databases.FindByIdsWhere(r.dB, roles, roleIds, wheres...)
	return roles, err
}

func (r *Role) FindWhere(query interface{}, wheres ...databases.Condition) (roles []*model.Role, err error) {
	err = databases.FindWhere(r.dB, roles, query, wheres...)
	return roles, err
}
