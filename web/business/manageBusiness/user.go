package manageBusiness

import (
	"app/logic"
	"app/model"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zvalid"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

// CurrentUser 当前登录用户
func CurrentUser(c *znet.Context) (user *model.AuthUser, has bool) {
	u, h := c.Value("user")
	if h {
		user, h = u.(*model.AuthUser)

	}
	return user, h
}

func UserOthersTokenDisable(tokenModel *model.AuthUserToken) {
	t := tokenModel.SelectUser()
	if flag := IsMultipleLogins(); !flag {
		t.UpdateUserToken()
	}
}

func Update(c *znet.Context, postData PutUpdateSt, currentUserId uint, editPwdAuth bool, isMe bool) (int64, error) {
	editUser := &model.AuthUser{ID: currentUserId}
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
		"email": valid.IsMail("email 地址错误").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if has, _ := editUser.EmailExist(postData.Email); has {
				return rawValue, errors.New("email 已被使用")
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

	updateUser := &model.AuthUser{
		Status:   postData.Status,
		Remark:   postData.Remark,
		Email:    postData.Email,
		Nickname: postData.Nickname,

		Password: postData.Password,
		GroupID:  postData.GroupID,
	}

	updateUser.Avatar, _ = MvAvatar(postData.Avatar, avatarFilename)
	queryFiled := []string{"update_time", "status", "avatar", "remark", "email", "nickname"}

	if editPwdAuth && postData.Password != "" {
		queryFiled = append(queryFiled, "password")
	}

	if editPwdAuth && !isMe {
		queryFiled = append(queryFiled, "group_id")
	}

	updateRes, err := (&model.AuthUser{}).Update(queryFiled, editUser, updateUser)
	if err != nil {
		return 0, errors.New("服务繁忙，请重试")
	}

	if err = (&model.AuthUserLogs{Userid: editUser.ID}).UpdatePasswordTip(c); err != nil {
		return 0, err
	}

	return updateRes, nil
}

type PutUpdateSt struct {
	Id        uint   `json:"id"`
	Status    uint8  `json:"status"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Avatar    string `json:"avatar"`
	Remark    string `json:"remark"`
	Email     string `json:"email"`
	GroupID   []uint `json:"group_id"`
	Nickname  string `json:"nickname"`
}

// 判断是否管理员, 上面还有一个超级管理员级别
func IsAdmin(userid uint) int {
	cfg := gconf.New("conf.yml")
	_ = cfg.Read()
	adminCfg := cfg.Get("admin")
	if adminCfg != nil {
		userSlice := strings.Split(adminCfg.(string), ",")
		for _, user := range userSlice {
			userStr := strings.TrimSpace(user)
			if userStr == ztype.ToString(userid) {
				return 1
			}
		}
	}

	pool := map[uint]int8{
		1: 1,
	}

	_, has := pool[userid]
	if has == true {
		return 1
	}

	return 0
}

// 移动临时头像文件到指定目录上
func MvAvatar(path string, filename string) (newPath string, err error) {
	if path == "" {
		return "", nil
	}
	path = avatarPrefix + path
	completePath := zfile.RealPathMkdir(avatarPath, true)

	newAva := completePath + filename

	if has := strings.HasPrefix(path, "/"); has {
		path = path[1:]
	}

	err = os.Rename(zfile.RealPath(path), newAva)
	if err != nil {
		return path, nil
	}

	rePath := "/" + zfile.SafePath(newAva)
	rePath = strings.TrimPrefix(rePath, avatarPrefix)

	return rePath, nil
}

// 图片上传
// path 相对路径+mad5(sha1(文件名))
func UploadImg(file *multipart.FileHeader, dist string, currentHost string) (path string, host string, err error) {
	fileName := file.Filename
	fileSuffix := fileName[strings.LastIndex(fileName, "."):]

	var inType = 0
	for _, v := range avatarType {
		if fileSuffix == v {
			inType = 1
			break
		}
	}
	if inType == 0 {
		return "", "", errors.New("文件类型错误！只允许：jpg | png")
	}

	fileSize := logic.Byte2Kb(file.Size)
	if fileSize > avatarSize {
		return "", "", errors.New(fmt.Sprintf("文件\"%v\"大小错误！最大：%vKB", fileName, avatarSize))
	}

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	src.Seek(0, 0)
	hasher := sha1.New()
	io.Copy(hasher, src)
	fileKey := hex.EncodeToString(hasher.Sum(nil))
	md5FileKey := fmt.Sprintf("%x", md5.Sum([]byte(fileKey)))

	path = dist + md5FileKey + fileSuffix
	out, err := os.Create(path)
	if err != nil {
		return "", "", err
	}
	defer out.Close()
	src.Seek(0, 0)
	_, err = io.Copy(out, src)
	if err != nil {
		return "", "", err
	}

	host = currentHost

	return "/" + zfile.SafePath(path), host, nil
}

// 上传用户头像
func UploadAvatar(file *multipart.FileHeader, currentHost string) (rePath string, host string, err error) {
	path := zfile.RealPathMkdir(avatarTemPath, true)
	rePath, host, err = UploadImg(file, path, currentHost)
	rePath = strings.TrimPrefix(rePath, avatarPrefix)

	return
}

// 判断当前token是否过期
func IsExpire(tokenInfo *model.AuthUserToken) error {
	tokenInfo.TokenValues()
	if LoginExpireTime() == 0 {
		return nil
	}

	h, _ := time.ParseDuration(fmt.Sprintf("%ds", LoginExpireTime()))
	lastTime, _ := ztime.Parse(ztime.FormatTime(tokenInfo.UpdatedAt.Time, "Y-m-d H:i:s"))
	nowTime, _ := ztime.Parse(ztime.Now("Y-m-d H:i:s"))
	if flag := nowTime.Before(lastTime.Add(1 * h)); !flag { // 接口有效时间
		tokenInfo.SetExpiration()
		return errors.New("登录过期，请重新登录")
	}

	return nil
}
