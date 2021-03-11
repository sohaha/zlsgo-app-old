package manage

import (
	"app/web/business/manageBusiness"
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
			"PUT": {
				"/ZlsManage/UserApi/EditPassword.go",
			},
			"POST": {
				"/ZlsManage/UserApi/GetToken.go",
				"/ZlsManage/UserApi/ClearToken.go",
			},
			"GET": {
				"/ZlsManage/UserApi/UnreadMessageCount.go",
				"/ZlsManage/UserApi/UseriInfo.go",
				"/ZlsManage/UserManageApi/UserLists.go",
				"/ZlsManage/UserApi/Logs.go",
			},
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
			token.Token = t
			deToken, err := token.TokenRules()
			if err != nil {
				c.ApiJSON(401, "请先登录", nil)
				return
			}

			if tokenId, err := strconv.Atoi(strings.Split(deToken, "|")[2]); err != nil {
				token.ID = uint(tokenId)
			}

			if err := manageBusiness.IsExpire(token); err != nil {
				c.ApiJSON(401, err.Error(), nil)
				return
			}

			if  token.Userid != 0 {
				token.UpdateTimeAndReturnUser(user)
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
		currentRule := &model.RuleCollation{
			AdoptRoute:     map[string][]string{},
			InterceptRoute: map[string][]string{},
			Marks:          []string{},
		}
		for _, groupID := range user.GroupID {
			currentRule = (&model.AuthUserGroup{Status: 1, ID: groupID}).GetRuleCollation(currentRule)
		}

		// currentRule := (&model.AuthUserGroup{Status: 1, ID: user.GroupID}).GetRuleCollation()
		if currentRule == nil {
			// 没有对应的权限规则，禁止访问
			c.ApiJSON(403, "请先为当前用户设置权限", nil)
			return
		}
		ruleMarks = currentRule.Marks

		for _, groupID := range user.GroupID {
			if groupID == GROUP_ADMIN && VerifRoutingPermission(path, method, &model.RuleCollation{
				AdoptRoute: map[string][]string{
					"GET": {
						"/ZlsManage/SystemApi/SystemConfig.go",
						"/ZlsManage/SystemApi/SystemLogs.go",
					},
					"POST": {
						"/ZlsManage/MenuApi/UserMenu.go",
					},
				},
			}) {
				c.Next()
				return
			}
		}

		if !VerifRoutingPermission(path, method, currentRule) {
			c.ApiJSON(403, "对不起，权限不足", nil)
			return
		}
		c.Next()
	}
}
