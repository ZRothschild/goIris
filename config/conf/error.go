package conf

import (
	"errors"
)

var (
	// 记录已存在
	ErrRecordExist = errors.New("已存在")
)
