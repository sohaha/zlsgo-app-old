package manage

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"strconv"
	"strings"

	"app/model"
)

const GROUP_ADMIN uint = 1

// VerifRoutingPermission 路由权限验证
func VerifRoutingPermission(currentPath, method string, rules *model.RuleCollation) (adopt bool) {
	routes := rules.AdoptRoute[method]
	for i := range routes {
		if zstring.Match(currentPath, routes[i]) {
			adopt = true
			break
		}

		if !strings.HasSuffix(routes[i], ".go") {
			if zstring.Match(currentPath, routes[i]+".go") {
				adopt = true
				break
			}
		}
	}

	if adopt {
		routes = rules.InterceptRoute[method]
		for i := range routes {
			if zstring.Match(currentPath, routes[i]) {
				return false
			}

			if !strings.HasSuffix(routes[i], ".go") {
				if zstring.Match(currentPath, routes[i]+".go") {
					return false
				}
			}
		}
	}
	return
}

// VerifPermissionMark 验证是否拥有指定权限标识码
func VerifPermissionMark(c *znet.Context, mark string, disable ...bool) (adopt bool) {
	user := &model.AuthUser{}
	if u, ok := c.Value("user"); ok {
		user = u.(*model.AuthUser)
	}
	// 超级管理员无需验证权限
	if user.IsSuper {
		return true
	}

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
			"POST": {"/ZlsManage/UserApi/GetToken.go", "/ZlsManage/UserApi/ClearToken.go"},
			"GET":  {"/ZlsManage/SystemApi/Logs.go", "/ZlsManage/SystemApi/UnreadMessageCount.go"},
		},
	}
	return func(c *znet.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path

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
			idx := strings.Index(t, "_")
			if idx >= 0 {
				tokenID, _ := strconv.Atoi(t[0:idx])

				token.ID = uint(tokenID)
				token.Token = t[idx+1:]
				if err := user.TokenToInfo(token); err != nil {
					c.ApiJSON(401, err.Error(), nil)
					return
				}
			}
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
