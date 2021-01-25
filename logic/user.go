package logic

import (
	"app/model"
	"app/web/business/manageBusiness"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zvalid"
)

func UserOthersTokenDisable(tokenModel *model.AuthUserToken) {
	t := tokenModel.SelectUser()
	if cfg, _ := (&manageBusiness.ParamPutSystemConfigSt{}).GetConf(); cfg.LoginMode {
		t.UpdateUserToken()
	}
}

func Update(c *znet.Context, postData manageBusiness.PutUpdateSt, currentUserId uint, editPwdAuth bool) (int64, error) {
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

	updateUser.Avatar, _ = manageBusiness.MvAvatar(postData.Avatar, avatarFilename)
	queryFiled := []string{"update_time", "status", "avatar", "remark", "email", "nickname"}

	if editPwdAuth && postData.Password != "" {
		queryFiled = append(queryFiled, "password")
	}

	// if isAdmin == 1 && !isMe {
	queryFiled = append(queryFiled, "group_id")
	// }

	updateRes, err := (&model.AuthUser{}).Update(queryFiled, editUser, updateUser)
	if err != nil {
		return 0, errors.New("服务繁忙，请重试")
	}

	if err = (&model.AuthUserLogs{Userid: editUser.ID}).UpdatePasswordTip(c); err != nil {
		return 0, err
	}

	return updateRes, nil
}
