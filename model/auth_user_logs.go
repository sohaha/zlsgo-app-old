package model

import (
	"gorm.io/gorm"
)

// AuthUserLogs 用户日志
type AuthUserLogs struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Userid    int            `gorm:"column:userid;type:int(11);not null;default:0;comment:对应用户Id;" json:"userid"`
	OperateID int            `gorm:"column:operate_id;type:int(11);not null;default:0;comment:操作人Id，游客为0;" json:"operate_id"`
	Title     string         `gorm:"column:title;type:varchar(100);not null;default:'';comment:标题;" json:"title"`
	Content   string         `gorm:"column:content;type:text(0);not null;comment:信息;" json:"content"`
	Type      uint8          `gorm:"column:type;type:tinyint(4);not null;default:1;comment:类型:1正常，2警告，3错误;" json:"type"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:1;comment:状态:1未读，2已读;" json:"status"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}
