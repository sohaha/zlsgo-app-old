package manage

import (
	"github.com/sohaha/zlsgo/znet"

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

		// 判断当前路由需要权限不
		currentRule, _ := rule[user.GroupID]
		if currentRule == nil {

		}
		c.Log.Dump(currentRule)
		c.Next()
	}
}
