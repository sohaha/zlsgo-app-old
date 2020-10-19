package model

import (
	"app/web/business/manageBusiness"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zvalid"
)

// AuthUser 管理员
type AuthUser struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Username  string         `gorm:"column:username;type:varchar(255);not null;default:'';comment:用户名;" json:"username"`
	Password  string         `gorm:"column:password;type:varchar(255);not null;default:'';comment:用户密码;" json:"-"`
	Key       string         `gorm:"column:key;type:varchar(255);not null;default:'';comment:密码盐;" json:"-"`
	Nickname  string         `gorm:"column:nickname;type:varchar(255);not null;default:'';comment:用户昵称;" json:"nickname"`
	Email     string         `gorm:"column:email;type:varchar(255);not null;default:'';comment:Email;" json:"email"`
	Remark    string         `gorm:"column:remark;type:varchar(255);not null;default:'';comment:用户简介;" json:"remark"`
	Avatar    string         `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:头像;" json:"avatar"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:0;comment:状态:-1软删除,0待激活,1正常,2禁止;" json:"status"`
	GroupID   uint           `gorm:"autoIncrement;column:group_id;type:int(11);not null;default:0;comment:角色Id;" json:"group_id"`
	IsSuper   bool           `gorm:"column:is_super;default:0;" json:"is_super"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

// 默认管理员密码
const DefManagePassword = "admin666"

func (u *AuthUser) Lists(pp *Page) (users []AuthUser) {
	_, _ = FindPage(context.Background(), db.Where(u).Model(u).Order("id desc"), pp, &users)
	return
}

func (*migrate) CreateAuthUser() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		password := DefManagePassword // zstring.Rand(6)
		encryptPassword, _ := zvalid.Text(password).EncryptPassword().String()
		return "CreateAuthUser", func(db *gorm.DB) error {
			db.Create([]AuthUser{
				{
					Username: "manage",
					Password: encryptPassword,
					Key:      "",
					Nickname: "超级管理员",
					Email:    "manage@qq.com",
					Avatar:   "",
					GroupID:  1,
					Status:   1,
					IsSuper:  true,
				},
				{
					Username: "admin",
					Password: encryptPassword,
					Key:      "",
					Nickname: "管理员",
					Email:    "admin@qq.com",
					Avatar:   "",
					GroupID:  1,
					Status:   1,
					IsSuper:  false,
				},
				{
					Username: "edit",
					Password: encryptPassword,
					Key:      "",
					Nickname: "编辑",
					Email:    "edit@qq.com",
					Avatar:   "",
					GroupID:  2,
					Status:   1,
					IsSuper:  false,
				},
			})
			return nil
		}
	})
}

func (u *AuthUser) Insert() error {
	has, err := u.UserExist()
	if err != nil {
		return err
	}
	if has {
		return errors.New("用户已存在")
	}
	tx := db.Create(u)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *AuthUser) UserExist() (bool, error) {
	where := AuthUser{}
	if u.Username != "" {
		where.Username = u.Username
	}
	if u.Email != "" {
		where.Email = u.Email
	}
	return Exist(context.Background(), db.Where("username = ?", where.Username).Or("email = ?", where.Email).Model(&where))
}

func (u *AuthUser) TokenToInfo(t *AuthUserToken) {
	t.Status = 1
	db.Where(t).Limit(1).Find(&t)
	if t.Userid != 0 {
		db.Where(&AuthUser{ID: t.Userid}).Limit(0).Find(&u)
	}
}

func (u *AuthUser) Login(ip string, ua string) (string, error) {
	password := u.Password
	db.Model(&u).Where(&AuthUser{Username: u.Username}).Limit(1).Find(&u)
	if u.ID == 0 {
		return "", errors.New("用户不存在")
	}
	if err := zvalid.Text(password, "用户密码").CheckPassword(u.Password).Error(); err != nil {
		return "", err
	}

	t := AuthUserToken{
		IP:     ip,
		UA:     ua,
		Userid: u.ID,
	}
	token := t.CreateToken()
	if token == "" {
		return "", errors.New("创建 Token 失败")
	}
	return t.Token, nil
}

func (u *AuthUser) GetUser() {
	u.Status = 1
	db.Where(u).Find(&u)
}

func (u *AuthUser) EmailExist(email string) (bool, error) {
	return Exist(context.Background(), db.Where("email = ? and id != ?", email, u.ID).Model(&AuthUser{}))
}

func (u *AuthUser) Update(c *znet.Context, postData manageBusiness.PutUpdateSt, currentUserId uint, isAdmin int, isMe bool) (int64, error) {
	editUser := &AuthUser{ID: currentUserId}
	(editUser).GetUser()

	valid := c.ValidRule()
	err := c.BindValid(&postData, map[string]zvalid.Engine{
		"remark": valid.MaxLength(200, "用户简介最多200字符"),
		"avatar": valid.MaxLength(250, "头像地址不能超过250字符"),
		"status": valid.EnumInt([]int{1, 2}, "用户状态值错误"),
		"password": valid.MinLength(3, "密码最少3字符").MaxLength(50, "密码最多50字符").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			if rawValue != postData.Password2 {
				newErr = errors.New("两次密码不一致")
			}
			newValue = rawValue
			return
		}).EncryptPassword(),
		"email": valid.IsMail("Email地址错误").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if has, _ := editUser.EmailExist(postData.Email); has {
				return rawValue, errors.New("Email已被使用")
			}
			newValue = rawValue
			return
		}),
	})
	if err != nil {
		return 0, err
	}

	idStr := ztype.ToString(currentUserId)
	md5Id := fmt.Sprintf("%x", md5.Sum([]byte(idStr)))
	avatarFilename := "user_" + md5Id + ".png"

	updateUser := &AuthUser{
		Status:   postData.Status,
		Remark:   postData.Remark,
		Email:    postData.Email,
		Nickname: postData.Nickname,

		Password: postData.Password,
		GroupID:  postData.GroupID,
	}

	updateUser.Avatar, _ = manageBusiness.MvAvatar(postData.Avatar, avatarFilename)
	queryFiled := []string{"update_time", "status", "avatar", "remark", "email", "nickname"}

	if isAdmin == 1 && postData.Password != "" {
		queryFiled = append(queryFiled, "password")
	}

	if isAdmin == 1 && !isMe {
		queryFiled = append(queryFiled, "group_id")
	}

	uRes := db.Model(&AuthUser{}).Select(queryFiled).Where("id = ?", editUser.ID).Updates(updateUser)
	if uRes.RowsAffected < 1 {
		return 0, errors.New("服务繁忙,请重试.")
	}

	if err = (&AuthUserLogs{Userid: editUser.ID}).UpdatePasswordTip(c); err != nil {
		return 0, err
	}

	return uRes.RowsAffected, nil
}

func (u *AuthUser) EditPassword(c *znet.Context, postData manageBusiness.PutEditPasswordSt) error {
	editUser := &AuthUser{ID: u.ID}
	(editUser).GetUser()
	if editUser.Email == "" {
		return errors.New("用户不存在")
	}

	if err := zvalid.Text(postData.OldPass, "原密码").CheckPassword(editUser.Password, "原密码错误").Error(); err != nil {
		return err
	}

	uRes := db.Model(&AuthUser{}).Select("password").Where("id = ?", editUser.ID).Updates(&AuthUser{Password: postData.Pass})
	if uRes.RowsAffected < 1 {
		return errors.New("服务繁忙,请重试.")
	}

	if err := (&AuthUserLogs{Userid: editUser.ID}).UpdatePasswordTip(c); err != nil {
		return err
	}

	return nil
}
