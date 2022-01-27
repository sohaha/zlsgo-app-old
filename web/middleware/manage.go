package middleware

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"app/conf"
	"app/dal/model"
	"app/dal/query"
	"app/global"
	"app/web"
	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/golang-jwt/jwt"
	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/zutil"
	"gorm.io/gorm"
)

var ManageCache = zcache.New("manage")

func Manage(pubRouters []string) func(c *znet.Context) {
	modelConf := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act, eft

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && keyMatch(r.act, p.act)
`
	dbConf := conf.DB()
	a, err := gormadapter.NewAdapterByDBUseTableName(global.DB, strings.TrimSuffix(dbConf.Prefix, "_"), "manage_permission")
	zutil.CheckErr(err, true)
	m, err := casbinModel.NewModelFromString(modelConf)
	zutil.CheckErr(err, true)
	e, err := casbin.NewEnforcer(m, a)
	zutil.CheckErr(err, true)
	zutil.CheckErr(e.LoadPolicy(), true)

	return func(c *znet.Context) {
		path := c.Request.URL.Path
		var ok bool
		for i := range pubRouters {
			if pubRouters[i] == path {
				ok = true
				break
			}
		}

		if !ok {
			token := getToken(c)
			if token == "" {
				web.ApiJSON(c, 401, "请先登录", nil)
				return
			}
			j, err := parsingToken(token)
			zlog.Debug(j, err)
			if err != nil || len(j.U) < 4 {
				web.ApiJSON(c, 401, "登录状态过期，请重新登录", err)
				return
			}

			key := j.U[:4]
			uid := j.U[4:]
			id, err := strconv.Atoi(uid)
			if err != nil {
				web.ApiJSON(c, 401, "无效用户", nil)
				return
			}
			u, err := ManageCache.MustGet(uid, func(set func(data interface{}, lifeSpan time.Duration, interval ...bool)) (err error) {
				u := query.Use(global.DB).ManageUser
				first, err := u.Where(u.ID.Eq(int32(id))).First()
				if err == nil {
					set(first, time.Minute*10, true)
				}
				return err
			})
			if err != nil {
				errMsg := "无效用户"
				if err == gorm.ErrRecordNotFound {
					errMsg = "用户不存在"
				}
				web.ApiJSON(c, 401, errMsg, nil)
				return
			}
			user, _ := u.(*model.ManageUser)
			if user.Status != 1 {
				web.ApiJSON(c, 403, "该账号已停用", nil)
				return
			}
			if user.Key != key {
				web.ApiJSON(c, 401, "登录状态失效，请重新登录", nil)
				return
			}
			now := time.Now().Unix()
			diff := j.ExpiresAt - now
			manageConf := conf.Manage()
			if diff <= int64(manageConf.Expire/3) {
				ResetToken(c, user.ID, user.Key)
			}
			if ok, err := e.Enforce(uid, path, c.Request.Method); !ok {
				web.ApiJSON(c, 403, "权限不足", err)
				return
			}
		}

		c.Next()
	}
}

type JwtInfo struct {
	U string `json:"u"`
	jwt.StandardClaims
}

func ResetToken(c *znet.Context, uid int32, key string) {
	token, err := CreateToken(uid, key)

	if err == nil {
		c.SetHeader("Re-Token", token)
	}
}

func CreateToken(uid int32, key string) (string, error) {
	u := key + strconv.Itoa(int(uid))
	expire := conf.Manage().Expire
	if expire == 0 {
		expire = 3600
	}
	var claims = JwtInfo{
		u,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expire) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(zstring.String2Bytes(conf.Manage().Key))
	if err != nil {
		return "", fmt.Errorf("生成token失败: %v", err)
	}
	return signedToken, nil
}

func parsingToken(token string) (*JwtInfo, error) {
	t, err := jwt.ParseWithClaims(token, &JwtInfo{}, func(token *jwt.Token) (i interface{}, err error) {
		return zstring.String2Bytes(conf.Manage().Key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*JwtInfo); ok && t.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}

func getToken(c *znet.Context) string {
	authorization := c.GetHeader("Authorization")
	if authorization != "" {
		authorization = authorization[len("Basic "):]
		split := strings.Split(authorization, ".")
		if len(split) == 3 {
			return authorization
		}
		v, err := zstring.Base64Decode(zstring.String2Bytes(authorization))
		if err != nil {
			return ""
		}
		return strings.Split(zstring.Bytes2String(v), ":")[0]
	}
	return c.DefaultFormOrQuery("Authorization", "")
}
