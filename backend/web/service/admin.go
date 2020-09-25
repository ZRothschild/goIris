package service

import (
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/logger"
	"sync"
)

var (
	onceAdmin = new(sync.Once)
)

type Admin struct {
	admin *repository.Admin
	log   *logger.Logger
}

// 单利admin service
func NewAdmin(admin *repository.Admin, log *logger.Logger) (adminSrv *Admin) {
	onceAdmin.Do(func() {
		adminSrv = &Admin{admin: admin, log: log}
	})
	return adminSrv
}
