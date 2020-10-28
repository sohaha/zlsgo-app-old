package manageBusiness

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/ztype"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type PutUpdateSt struct {
	Id        uint   `json:"id"`
	Status    uint8  `json:"status"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Avatar    string `json:"avatar"`
	Remark    string `json:"remark"`
	Email     string `json:"email"`
	GroupID   uint   `json:"group_id"`
	Nickname  string `json:"nickname"`
}

type PutEditPasswordSt struct {
	OldPass string `json:"oldPass"`
	Pass    string `json:"pass"`
	Pass2   string `json:"pass2"`
	UserID  uint   `json:"userid"`
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
	path = avatarPrefix + path
	completePath := zfile.RealPathMkdir(avatarPath, true)

	newAva := completePath + filename

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

	fileSize := Byte2Kb(file.Size)
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
