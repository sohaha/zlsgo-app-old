package model

import (
	"app/web/business/manageBusiness"
	"context"
	"crypto/md5"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zvalid"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type GroupIdArr []uint

// AuthUser 管理员
//GroupID   uint           `gorm:"column:group_id;type:int(11);not null;default:0;comment:角色Id;" json:"group_id"`
//GroupID2  GroupIdArr     `gorm:"column:group_id2;type:varchar(255);not null;default:'';comment:角色Id2;" json:"group_id2"`
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
		fmt.Println(i)
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
	/*if string(val) == "null" {
		return nil
	}*/

	arrStr := strings.Split(string(val), ",")
	arrUint := []uint{}
	for _, str := range arrStr {
		if parInt, err := strconv.Atoi(str); err == nil { // 因为这里数据不应该存在中文等字符
			arrUint = append(arrUint, uint(parInt))
		}
	}

	//*g = GroupIdArr(arrUint)
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
	groups := []AuthUserGroup{}
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
			db.Create([]AuthUser{
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
			})
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

	h, _ := time.ParseDuration(TOKEN_EFFECTIVE_TIME)
	lastTime, _ := ztime.Parse(ztime.FormatTime(t.UpdatedAt.Time, "Y-m-d H:i:s"))
	nowTime, _ := ztime.Parse(ztime.Now("Y-m-d H:i:s"))
	if flag := nowTime.Before(lastTime.Add(1 * h)); !flag { // 接口有效时间
		db.Model(&t).Select("status", "update_time").Updates(AuthUserToken{Status: TOKEN_DISABLED}) // 让token过期
		return errors.New("登录过期，请重新登录")
	}

	if t.Userid != 0 {
		db.Model(&t).Select("update_time").Updates(AuthUserToken{}) // 更新token时间
		if cfg, _ := (&manageBusiness.ParamPutSystemConfigSt{}).GetConf(); cfg.LoginMode {
			db.Model(&AuthUserToken{}).Select("update_time, status").Where("userid = ? and id != ? and status != ?", t.Userid, t.ID, TOKEN_DISABLED).Updates(AuthUserToken{Status: TOKEN_DISABLED})
		}

		db.Where(&AuthUser{ID: t.Userid}).Limit(0).Find(&u)
	}

	return nil
}

func (u *AuthUser) Login(ip string, ua string) (string, uint, error) {
	password := u.Password
	db.Model(&u).Where(&AuthUser{Username: u.Username}).Limit(1).Find(&u)
	if u.ID == 0 {
		return "", 0, errors.New("用户不存在")
	}
	if err := zvalid.Text(password, "用户密码").CheckPassword(u.Password).Error(); err != nil {
		return "", 0, err
	}

	t := AuthUserToken{
		IP:     ip,
		UA:     ua,
		Userid: u.ID,
	}
	token := t.CreateToken()
	if token == "" {
		return "", 0, errors.New("创建 Token 失败")
	}
	return t.Token, t.ID, nil
}

func (u *AuthUser) GetUser() {
	u.Status = 1
	db.Where(u).Find(&u)
}

func (u *AuthUser) EmailExist(email string) (bool, error) {
	return Exist(context.Background(), db.Where("email = ? and id != ?", email, u.ID).Model(&AuthUser{}))
}

func (u *AuthUser) Update(c *znet.Context, postData manageBusiness.PutUpdateSt, currentUserId uint, editPwdAuth bool) (int64, error) {
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

	if editPwdAuth && postData.Password != "" {
		queryFiled = append(queryFiled, "password")
	}

	// if isAdmin == 1 && !isMe {
	queryFiled = append(queryFiled, "group_id")
	// }

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

func (u *AuthUser) Delete() error {
	var count int64
	if db.Model(&AuthUser{}).Where("id > ?", 0).Count(&count); count <= 1 {
		return errors.New("不予许删除唯一用户")
	}

	res := db.Delete(&AuthUser{}, u.ID)
	if res.RowsAffected < 1 {
		return errors.New("服务繁忙,请重试.")
	}

	return nil
}
