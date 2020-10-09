package manage

import (
	"strings"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"

	"app/model"
)

// VerifRoutingPermission 路由权限验证
func VerifRoutingPermission(currentPath, method string, rules *model.RuleCollation) (adopt bool) {
	routes := rules.AdoptRoute[method]
	for i := range routes {
		if zstring.Match(currentPath, routes[i]) {
			adopt = true
			break
		}
	}

	if adopt {
		routes = rules.InterceptRoute[method]
		for i := range routes {
			if zstring.Match(currentPath, routes[i]) {
				return false
			}
		}
	}
	return
}

// VerifPermissionMark 验证是否拥有指定权限标识码
func VerifPermissionMark(c *znet.Context, mark string, disable ...bool) (adopt bool) {
	r, ok := c.Value("ruleMarks")
	if ok {
		ruleMarks := *r.(*[]string)
		for _, v := range ruleMarks {
			if v == mark {
				return true
			}
		}
	}

	if !adopt {
		if len(disable) > 0 && disable[0] {
			return
		}
		c.ApiJSON(403, "对不起，无操作权限", nil)
	}
	return
}

func Authority() func(c *znet.Context) {
	ignoreRules := &model.RuleCollation{
		AdoptRoute: map[string][]string{
			"POST": {"/ZlsManage/UserApi/GetToken.go"},
		},
	}
	return func(c *znet.Context) {
		method := c.Request.Method
		path := strings.TrimLeft(c.Request.URL.Path, "/")

		user := &model.AuthUser{}
		token := &model.AuthUserToken{}
		var ruleMarks []string
		c.WithValue("user", user)
		c.WithValue("token", token)
		c.WithValue("ruleMarks", &ruleMarks)

		t := c.GetHeader("token")
		if t == "" {
			t = c.DefaultFormOrQuery("token", "")
		}
		if t != "" {
			token.Token = t
			user.TokenToInfo(token)
			// 后期可以考虑把用户信息缓存
		}

		if VerifRoutingPermission(path, method, ignoreRules) {
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

		// 获取角色的规则
		currentRule := (&model.AuthUserGroup{Status: 1, ID: user.GroupID}).GetRuleCollation()
		if currentRule == nil {
			// 没有对应的权限规则，禁止访问
			c.ApiJSON(403, "请先为当前用户设置权限", nil)
			return
		}
		ruleMarks = currentRule.Marks
		if !VerifRoutingPermission(path, method, currentRule) {
			c.ApiJSON(403, "对不起，权限不足", nil)
			return
		}
		c.Next()
	}
}
