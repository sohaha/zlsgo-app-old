package manage

import (
	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"gorm.io/gorm"

	"app/model"
)

var (
	rule   = map[uint]*[]model.AuthUserRules{}
	groups []model.AuthUserGroup
)

func getRule() {
	(&model.AuthUserGroup{Status: 1}).All(&groups)
	for _, v := range groups {
		rules := v.GetRules()
		rule[v.ID] = &rules
	}
}

var ruleCache = zcache.New("ruleCache")

func Authority() func(c *znet.Context) {
	go getRule()

	return func(c *znet.Context) {
		user := &model.AuthUser{}
		token := c.GetHeader("token")
		if token == "" {
			token = c.DefaultFormOrQuery("token", "")
		}

		if token != "" {
			user.TokenToInfo(token)
		}

		currentRuleArray := (&model.AuthUserGroup{Status: 1, Model: gorm.Model{ID: user.GroupID}}).GetRules()
		if len(currentRuleArray) == 0 {
			// 没有对应的权限规则，禁止访问
			c.ApiJSON(401, "请先登录", nil)
			return
		}

		path := c.Request.URL.Path
		adopt := false
		// 判断当前路由需要权限不
		for _, v := range currentRuleArray {
			if v.Type == 1 {
				if adopt {
					continue
				}
				// 路由类型 进行模糊匹配
				if zstring.Match(path, v.Mark) {
					adopt = true
				}
			} else {
				// 关键字类型
			}
		}

		if !adopt {
			c.ApiJSON(403, "对不起，权限不足", nil)
			return
		}
		c.Next()
	}
}
