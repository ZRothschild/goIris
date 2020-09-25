package service

import (
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/logger"
	"sync"
)

var (
	oncePermission = new(sync.Once)
)

type Permission struct {
	permission *repository.Permission
	log        *logger.Logger
}

// 单利permission service
func NewPermission(permission *repository.Permission, log *logger.Logger) (permissionSrv *Permission) {
	oncePermission.Do(func() {
		permissionSrv = &Permission{permission: permission, log: log}
	})
	return permissionSrv
}
