package repository

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/utils/lib/databases"
	"gorm.io/gorm"
)

// 用户
type Admin struct {
	admin *model.Admin
	dB    *gorm.DB
}

// 用户示例
func NewAdmin(admin *model.Admin, dB *gorm.DB) *Admin {
	return &Admin{
		admin: admin,
		dB:    dB,
	}
}

// 创建
func (r *Admin) Creat(admin *model.Admin) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Create(r.dB, admin)
	return
}

// Save 会保存所有的字段，即使字段是零值
func (r *Admin) Save(admin *model.Admin) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Save(r.dB, admin)
	return
}

// 修改
func (r *Admin) Updates(admin *model.Admin, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Updates(r.dB, r.admin, admin, query, args...)
	return
}

// 删除
func (r *Admin) Delete(admin *model.Admin, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = databases.Delete(r.dB, admin, query, args...)
	return
}

func (r *Admin) First(where interface{}, args ...interface{}) (admin *model.Admin, err error) {
	err = databases.First(r.dB, admin, where, args...)
	return admin, err
}

func (r *Admin) FirstById(adminId uint64) (admin *model.Admin, err error) {
	err = databases.FirstById(r.dB, admin, adminId)
	return admin, err
}

func (r *Admin) FirstByIdWhere(adminId uint64, wheres ...databases.Condition) (admin *model.Admin, err error) {
	err = databases.FirstByIdWhere(r.dB, admin, adminId, wheres...)
	return admin, err
}

// 提交查询
func (r *Admin) FirstWhere(query interface{}, wheres ...databases.Condition) (admin *model.Admin, err error) {
	err = databases.FirstWhere(r.dB, admin, query, wheres...)
	return admin, err
}

func (r *Admin) Find(query interface{}, args ...interface{}) (admins []*model.Admin, err error) {
	err = databases.Find(r.dB, admins, query, args...)
	return admins, err
}

func (r *Admin) FindByIds(adminIds []uint64) ([]*model.Admin, error) {
	var (
		err    error
		admins = make([]*model.Admin, 0, len(adminIds))
	)
	err = databases.FindByIds(r.dB, admins, adminIds)
	return admins, err
}

func (r *Admin) FindByIdsWhere(adminIds []uint64, wheres ...databases.Condition) ([]*model.Admin, error) {
	var (
		err    error
		admins = make([]*model.Admin, 0, len(adminIds))
	)
	err = databases.FindByIdsWhere(r.dB, admins, adminIds, wheres...)
	return admins, err
}

func (r *Admin) FindWhere(query interface{}, wheres ...databases.Condition) (admins []*model.Admin, err error) {
	err = databases.FindWhere(r.dB, admins, query, wheres...)
	return admins, err
}
