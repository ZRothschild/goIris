package model

import (
	"github.com/ZRothschild/goIris/app/model/base"
)

type User struct {
	base.Base
	Name string
}

func NewUser() *User {
	return &User{}
}
