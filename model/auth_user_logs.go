package model

import (
	"gorm.io/gorm"
)

// AuthUserLogs 管理员日志
type AuthUserLogs struct {
	gorm.Model
	Userid    int
	OperateID int
	Content   string
	Title     string `gorm:"type:varbinary(220)"`
	Type      uint8
	Status    uint8 `gorm:"type:int(2);default:1"`
}
