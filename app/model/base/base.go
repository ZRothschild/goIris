package base

import (
	"time"
)

type Base struct {
	ID        uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT:false;" json:"id,string" form:"id"`    // 主键
	CreatedAt time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt time.Time `gorm:"index;comment:删除时间" json:"deletedAt" form:"deletedAt"`                      // 删除时间
}
