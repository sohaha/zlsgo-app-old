package model

import (
	"errors"
	"fmt"
	"github.com/sohaha/zlsgo/znet"
	"gorm.io/gorm"
)

const (
	LOG_TYPE_NORMAL = 1
	LOG_TYPE_WARN   = 2
	LOG_TYPE_ERROR  = 3
	LOG_STATUS_NOT  = 1
	LOG_STATUS_READ = 2
)

// AuthUserLogs 用户日志
type AuthUserLogs struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Userid    uint           `gorm:"column:userid;type:int(11);not null;default:0;comment:对应用户Id;" json:"userid"`
	OperateID uint           `gorm:"column:operate_id;type:int(11);not null;default:0;comment:操作人Id，游客为0;" json:"operate_id"`
	Title     string         `gorm:"column:title;type:varchar(100);not null;default:'';comment:标题;" json:"title"`
	Content   string         `gorm:"column:content;type:text(0);not null;comment:信息;" json:"content"`
	Type      uint8          `gorm:"column:type;type:tinyint(4);not null;default:1;comment:类型:1正常，2警告，3错误;" json:"type"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:1;comment:状态:1未读，2已读;" json:"status"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

func (c *AuthUserLogs) UpdatePasswordTip(ctx *znet.Context) (int64, error) {
	// 非自己修改了用户密码之后需要记录日志 修改密码之后要重置用户token
	fmt.Println(11111111111)
	user, _ := ctx.Value("user")

	/*var (
		tip    string
		level  int
		status int
	)*/
	c.OperateID = user.(*AuthUser).ID

	if c.Userid == user.(*AuthUser).ID {
		/*tip = "修改密码成功"
		level = LOG_TYPE_NORMAL
		status = LOG_STATUS_READ*/

		c.Content = "修改密码成功"
		c.Type = LOG_TYPE_NORMAL
		c.Status = LOG_STATUS_READ
	} else {
		/*tip = fmt.Sprintf("您的密码被[%v]修改!", user.(*AuthUser).Username)
		level = LOG_TYPE_WARN
		status = LOG_TYPE_NORMAL*/

		c.Content = fmt.Sprintf("您的密码被[%v]修改!", user.(*AuthUser).Username)
		c.Type = LOG_TYPE_WARN
		c.Status = LOG_TYPE_NORMAL
	}

	c.title(ctx)
	res := db.Select([]string{}).Create(&c)
	if res.RowsAffected < 1 {
		return 0, errors.New("服务繁忙,请重试.")
	}

	fmt.Println(user.(*AuthUser))

	return res.RowsAffected, nil
}

func (c *AuthUserLogs) title(ctx *znet.Context) {
	if c.Title == "" {
		cLen := len([]rune(c.Content))
		if cLen > 50 {
			cLen = 50
		}
		c.Title = string([]rune(c.Content)[:cLen])
	}
	c.Content = fmt.Sprintf("%v\nOperate IP: %v\nUser Agent: %v", c.Content, ctx.GetClientIP(), ctx.GetHeader("User-Agent"))
}
