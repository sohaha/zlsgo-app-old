package model

import (
	"context"
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/zvalid"
	"gorm.io/gorm"
)

type GroupIdArr []uint

// AuthUser 管理员
// GroupID   uint           `gorm:"column:group_id;type:int(11);not null;default:0;comment:角色Id;" json:"group_id"`
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
	GroupID   GroupIdArr     `gorm:"column:group_id;type:varchar(255);not null;default:'';comment:角色Id;" json:"group_id"`
	IsSuper   bool           `gorm:"column:is_super;default:0;" json:"is_super"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

func (g GroupIdArr) Value() (driver.Value, error) {
	sqlStr := zstring.Buffer()
	gLen := len(g)
	for i, gid := range g {
		if gLen-1 == i {
			sqlStr.WriteString(strconv.Itoa(int(gid)))
		} else {
			sqlStr.WriteString(strconv.Itoa(int(gid)) + ",")
		}
	}

	return sqlStr.String(), nil
}

func (g *GroupIdArr) Scan(v interface{}) error {
	val, _ := v.([]byte)
	arrStr := strings.Split(string(val), ",")
	var arrUint []uint
	for _, str := range arrStr {
		if parInt, err := strconv.Atoi(str); err == nil { // 因为这里数据不应该存在中文等字符
			arrUint = append(arrUint, uint(parInt))
		}
	}

	*g = arrUint

	return nil
}

// 默认管理员密码
const DefManagePassword = "admin666"

type ListsModel struct {
	AuthUser
	Groups []string `json:"groups"`
}

func (u *AuthUser) Lists(pp *Page) (users []AuthUser) {
	wCond := " 1 = 1"
	wParams := make([]interface{}, 0)
	if u.ID > 0 {
		wCond += " and id = ?"
		wParams = append(wParams, u.ID)
	}
	if u.Username != "" {
		wCond += " and username like ?"
		wParams = append(wParams, "%"+u.Username+"%")
	}
	_, _ = FindPage(context.Background(), db.Where(wCond, wParams...).Model(u).Order("id desc"), pp, &users)
	return
}

func (u *AuthUser) ListsSub(users []AuthUser) (lists []ListsModel) {
	groups := make([]AuthUserGroup, 0)
	(&AuthUserGroup{}).All(&groups)
	kV := map[uint]string{}
	for _, v := range groups {
		kV[v.ID] = v.Name
	}

	getGroups := func(user *AuthUser, pools map[uint]string) (re []string) {
		for _, groupID := range user.GroupID {
			if v, flag := pools[groupID]; flag {
				re = append(re, v)
			}
		}

		if user.IsSuper {
			re = append(re, "超级管理员")
		}

		return
	}

	for _, user := range users {
		lists = append(lists, ListsModel{
			AuthUser: user,
			Groups:   getGroups(&user, kV),
		})
	}

	return lists
}

func (*migrate) CreateAuthUser() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		password := DefManagePassword // zstring.Rand(6)
		encryptPassword, _ := zvalid.Text(password).EncryptPassword().String()
		return "CreateAuthUser", func(db *gorm.DB) error {
			builtInUsers := []AuthUser{
				{
					Username: "manage",
					Password: encryptPassword,
					Key:      "",
					Nickname: "超级管理员",
					Email:    "manage@qq.com",
					Avatar:   "",
					GroupID:  []uint{1},
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
					GroupID:  []uint{1},
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
					GroupID:  []uint{2},
					Status:   1,
					IsSuper:  false,
				},
			}
			tx := db.Create(builtInUsers)
			if tx.Error != nil {
				log.Fatalf("初始化内置用户失败: %s", tx.Error.Error())
			}
			log.Tips("初始化内置用户:")
			for _, v := range builtInUsers {
				log.Printf("      账号: %s 密码: %s", v.Username, password)
			}
			return nil
		}
	})
}

func (u *AuthUser) Insert(pwd string) error {
	has, err := u.UserExist()
	if err != nil {
		return err
	}
	if has {
		return errors.New("用户已存在")
	}
	u.Password = pwd
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

func (u *AuthUser) TokenToInfo(t *AuthUserToken) error {
	t.Status = 1
	db.Where(t).Limit(1).Find(&t)

	h, _ := time.ParseDuration(TokenEffectiveTime)
	lastTime, _ := ztime.Parse(ztime.FormatTime(t.UpdatedAt.Time, "Y-m-d H:i:s"))
	nowTime, _ := ztime.Parse(ztime.Now("Y-m-d H:i:s"))
	if flag := nowTime.Before(lastTime.Add(1 * h)); !flag { // 接口有效时间
		db.Model(&t).Select("status", "update_time").Updates(AuthUserToken{Status: TokenDisabled}) // 让token过期
		return errors.New("登录过期，请重新登录")
	}

	if t.Userid != 0 {
		db.Model(&t).Select("update_time").Updates(AuthUserToken{}) // 更新token时间

		db.Where(&AuthUser{ID: t.Userid}).Limit(0).Find(&u)
	}

	return nil
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

func (u *AuthUser) Update(queryFiled []string, editUser *AuthUser, updateUser *AuthUser) (int64, error) {
	uRes := db.Model(&AuthUser{}).Select(queryFiled).Where("id = ?", editUser.ID).Updates(updateUser)
	if uRes.RowsAffected < 1 {
		return 0, errors.New("服务繁忙，请重试")
	}

	return uRes.RowsAffected, nil
}

type PutEditPasswordSt struct {
	OldPass string `json:"oldPass"`
	Pass    string `json:"pass"`
	Pass2   string `json:"pass2"`
	UserID  uint   `json:"userid"`
}

func (u *AuthUser) EditPassword(c *znet.Context, oldPass string, newPass string) error {
	editUser := &AuthUser{ID: u.ID}
	(editUser).GetUser()
	if editUser.Email == "" {
		return errors.New("用户不存在")
	}

	if err := zvalid.Text(oldPass, "原密码").CheckPassword(editUser.Password, "原密码错误").Error(); err != nil {
		return err
	}

	uRes := db.Model(&AuthUser{}).Select("password").Where("id = ?", editUser.ID).Updates(&AuthUser{Password: newPass})
	if uRes.RowsAffected < 1 {
		return errors.New("服务繁忙，请重试")
	}

	if err := (&AuthUserLogs{Userid: editUser.ID}).UpdatePasswordTip(c); err != nil {
		return err
	}

	return nil
}

func (u *AuthUser) Delete() error {
	var count int64
	if db.Model(&AuthUser{}).Where("id > ?", 0).Count(&count); count <= 1 {
		return errors.New("不予许删除唯一用户")
	}

	res := db.Delete(&AuthUser{}, u.ID)
	if res.RowsAffected < 1 {
		return errors.New("服务繁忙，请重试")
	}

	return nil
}
