package repository

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/utils/lib/databases"
	"gorm.io/gorm"
)

// 用户
type Permission struct {
	permission *model.Permission
	dB         *gorm.DB
}

// 用户示例
func NewPermission(permission *model.Permission, dB *gorm.DB) *Permission {
	return &Permission{
		permission: permission,
		dB:         dB,
	}
}

// 创建
func (r *Permission) Creat(permission *model.Permission) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Create(r.dB, permission)
	return
}

// Save 会保存所有的字段，即使字段是零值
func (r *Permission) Save(permission *model.Permission) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Save(r.dB, permission)
	return
}

// 修改
func (r *Permission) Updates(permission *model.Permission, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Updates(r.dB, r.permission, permission, query, args...)
	return
}

// 删除
func (r *Permission) Delete(permission *model.Permission, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Delete(r.dB, permission, query, args...)
	return
}

func (r *Permission) First(where interface{}, args ...interface{}) (permission *model.Permission, err error) {
	err = databases.First(r.dB, permission, where, args...)
	return permission, err
}

func (r *Permission) FirstById(permissionId uint64) (permission *model.Permission, err error) {
	err = databases.FirstById(r.dB, permission, permissionId)
	return permission, err
}

func (r *Permission) FirstByIdWhere(permissionId uint64, wheres ...databases.Condition) (permission *model.Permission, err error) {
	err = databases.FirstByIdWhere(r.dB, permission, permissionId, wheres...)
	return permission, err
}

// 提交查询
func (r *Permission) FirstWhere(query interface{}, wheres ...databases.Condition) (permission *model.Permission, err error) {
	err = databases.FirstWhere(r.dB, permission, query, wheres...)
	return permission, err
}

func (r *Permission) Find(query interface{}, args ...interface{}) (permissions []*model.Permission, err error) {
	err = databases.Find(r.dB, permissions, query, args...)
	return permissions, err
}

func (r *Permission) FindByIds(permissionIds []uint64) ([]*model.Permission, error) {
	var (
		err         error
		permissions = make([]*model.Permission, 0, len(permissionIds))
	)
	err = databases.FindByIds(r.dB, permissions, permissionIds)
	return permissions, err
}

func (r *Permission) FindByIdsWhere(permissionIds []uint64, wheres ...databases.Condition) ([]*model.Permission, error) {
	var (
		err         error
		permissions = make([]*model.Permission, 0, len(permissionIds))
	)
	err = databases.FindByIdsWhere(r.dB, permissions, permissionIds, wheres...)
	return permissions, err
}

func (r *Permission) FindWhere(query interface{}, wheres ...databases.Condition) (permissions []*model.Permission, err error) {
	err = databases.FindWhere(r.dB, permissions, query, wheres...)
	return permissions, err
}
