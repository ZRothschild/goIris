package conf

import (
	"errors"
)

var (
	// 记录已存在
	ErrRecordExist = errors.New("record exist")
)
