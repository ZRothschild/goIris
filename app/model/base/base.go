package base

import (
	"time"
)

type Base struct {
	ID        uint64     `gorm:"column:id;primary_key;AUTO_INCREMENT:false;" json:"id,string" form:"id"`                  // 主键
	CreatedAt int64      `gorm:"column:created_at;not null;default:'0';comment:'创建时间'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt int64      `gorm:"column:updated_at;not null;default:'0';comment:'更新时间'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt *time.Time `gorm:"index;null;default:null;comment:'删除时间'" json:"deletedAt" form:"deletedAt"`                //删除时间
}
