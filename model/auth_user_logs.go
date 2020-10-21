package model

import (
	"context"
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

func (c *AuthUserLogs) UpdatePasswordTip(ctx *znet.Context) error {
	// 非自己修改了用户密码之后需要记录日志 修改密码之后要重置用户token
	user, _ := ctx.Value("user")
	c.OperateID = user.(*AuthUser).ID

	if c.Userid == user.(*AuthUser).ID {
		c.Content = "修改密码成功"
		c.Type = LOG_TYPE_NORMAL
		c.Status = LOG_STATUS_READ
	} else {
		c.Content = fmt.Sprintf("您的密码被[%v]修改!", user.(*AuthUser).Username)
		c.Type = LOG_TYPE_WARN
		c.Status = LOG_TYPE_NORMAL
	}

	c.title(ctx)
	res := db.Select([]string{}).Create(&c)
	if res.RowsAffected < 1 {
		return errors.New("服务繁忙,请重试.")
	}

	return nil
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

type LogListsModel struct {
	AuthUserLogs
	Username string `json:"username"`
}

func (c *AuthUserLogs) Lists(pp *Page) (logs []LogListsModel) {
	wCond := " 1 = 1"
	wParams := make([]interface{}, 0)
	wCond += " and auth_user_logs.`userid` = ?"
	wParams = append(wParams, c.Userid)
	if c.Type > 0 {
		wCond += " and auth_user_logs.`type` = ?"
		wParams = append(wParams, c.Type)
	}
	if c.Status > 0 {
		wCond += " and auth_user_logs.`status` = ?"
		wParams = append(wParams, LOG_STATUS_NOT)
	}

	_, _ = FindPage(context.Background(), db.Model(c).Select("auth_user_logs.*", "auth_user.username as username").Where(wCond, wParams...).Joins("LEFT JOIN auth_user ON auth_user.id = auth_user_logs.operate_id").Order("auth_user_logs.id desc"), pp, &logs)
	return
}

func (c *AuthUserLogs) UnreadMessageCount() (count int64) {
	db.Model(&AuthUserLogs{}).Where("id > ? and userid = ? and status = ?", c.ID, c.Userid, LOG_STATUS_NOT).Count(&count)
	return
}

func (c *AuthUserLogs) UpdateMessageStatus(ids []int) uint {
	c.Status = LOG_STATUS_READ
	res := db.Model(&AuthUserLogs{}).Select([]string{"update_time", "status"}).Where("id IN ? and userid = ?", ids, c.Userid).Updates(c)

	return uint(res.RowsAffected)
}
