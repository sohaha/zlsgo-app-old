package model

import (
	"gorm.io/gorm"
)

// AuthUserLogs 管理员日志
type AuthUserLogs struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt JSONTime       `gorm:"column:create_time;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Userid    int
	OperateID int
	Content   string
	Title     string `gorm:"type:varbinary(220)"`
	Type      uint8
	Status    uint8 `gorm:"type:int(2);default:1"`
}
