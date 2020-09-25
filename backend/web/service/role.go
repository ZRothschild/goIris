package service

import (
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/logger"
	"sync"
)

var (
	onceRole = new(sync.Once)
)

type Role struct {
	role *repository.Role
	log  *logger.Logger
}

// 单利role service
func NewRole(role *repository.Role, log *logger.Logger) (roleSrv *Role) {
	onceRole.Do(func() {
		roleSrv = &Role{role: role, log: log}
	})
	return roleSrv
}
