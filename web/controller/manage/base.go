package manage

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"

	"app/model"
)

var (
	rule   = map[uint]*[]model.AuthUserRules{}
	groups []model.AuthUserGroup
)

func verifRule(currentPath string, rules []model.AuthUserRules) (adopt bool) {
	// 判断当前路由需要权限不
	for _, v := range rules {
		if v.Type == 1 {
			if adopt {
				continue
			}
			// 路由类型 进行模糊匹配
			if zstring.Match(currentPath, v.Mark) {
				adopt = true
			}
		} else {
			// 关键字类型
		}
	}
	return
}

func Authority() func(c *znet.Context) {
	ignoreRules := []model.AuthUserRules{
		{
			Type: 1,
			Mark: "/ZlsManage/UserApi/GetToken.go",
		},
	}
	return func(c *znet.Context) {
		user := &model.AuthUser{}
		token := &model.AuthUserToken{}
		c.WithValue("user", user)
		c.WithValue("token", token)

		path := c.Request.URL.Path
		t := c.GetHeader("token")
		if t == "" {
			t = c.DefaultFormOrQuery("token", "")
		}

		if t != "" {
			token.Token = t
			user.TokenToInfo(token)
			// 后期可以考虑把用户信息缓存
		}

		if verifRule(path, ignoreRules) {
			c.Next()
			return
		}

		if user.ID == 0 {
			c.ApiJSON(401, "请先登录", nil)
			return
		}

		// 超级管理员无需验证权限
		if user.IsSuper {
			c.Next()
			return
		}

		currentRuleArray := (&model.AuthUserGroup{Status: 1, ID: user.GroupID}).GetRules()
		if len(currentRuleArray) == 0 {
			// 没有对应的权限规则，禁止访问
			c.ApiJSON(401, "请先当前用户设置权限", nil)
			return
		}

		if !verifRule(path, currentRuleArray) {
			c.ApiJSON(403, "对不起，权限不足", nil)
			return
		}

		c.Next()
	}
}
