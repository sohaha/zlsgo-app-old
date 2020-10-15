package model

import (
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zcache"
)

const (
	TYPE_ROUTER uint8 = 1 // 路由
	TYPE_MARK   uint8 = 2 // 标识码
)

// AuthUserGroup 用户角色
type AuthUserGroup struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Name      string         `gorm:"column:name;type:varchar(255);not null;default:'';comment:角色名称;" json:"name"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:1;comment:状态:1正常，2禁止;" json:"status"`
	Remark    string         `gorm:"column:remark;type:varchar(255);not null;default:'';comment:角色简介;" json:"remark"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

type (
	RuleCollation struct {
		AdoptRoute     map[string][]string
		InterceptRoute map[string][]string
		Marks          []string
	}
)

// All 获取全部记录
func (g AuthUserGroup) All(groups *[]AuthUserGroup) *gorm.DB {
	return db.Where(&AuthUserGroup{Status: g.Status}).Find(&groups)
}

var ruleCache = zcache.New("ruleCache")

// GetRules 获取规则列表
func (g AuthUserGroup) GetRules() (rules []AuthUserRules) {
	currentRule, err := ruleCache.MustGet(strconv.Itoa(int(g.ID)), func(set func(data interface{}, lifeSpan time.Duration, interval ...bool)) (err error) {
		var relas []AuthUserRulesRela
		db.Model(&AuthUserRulesRela{}).Where(map[string]interface{}{"group_id": g.ID, "status": 1}).Select("rule_id").Find(&relas)
		var ids []uint

		for _, v := range relas {
			ids = append(ids, v.RuleID)
		}
		db.Model(&AuthUserRules{}).Where("id in (?)", ids).Find(&rules)

		set(rules, 60*10, true)
		return nil
	})
	if err != nil {
		return
	}
	return currentRule.([]AuthUserRules)
}

// GetRuleCollation 获取整理后的规则列表
func (g AuthUserGroup) GetRuleCollation() (s *RuleCollation) {
	rules := g.GetRules()
	if len(rules) == 0 {
		return
	}
	// 有必要可以升级成树
	s = &RuleCollation{
		AdoptRoute:     map[string][]string{},
		InterceptRoute: map[string][]string{},
		Marks:          []string{},
	}
	setData := func(methods []string, route string, v AuthUserRules) {
		for _, m := range methods {
			if v.Status == AuthUserRulesStatusIntercept {
				s.InterceptRoute[m] = append(s.InterceptRoute[m], route)
			} else if v.Status == AuthUserRulesStatusAdopt {
				s.AdoptRoute[m] = append(s.AdoptRoute[m], route)
			}
		}
	}
	// 判断当前路由权限
	for _, v := range rules {
		if v.Type == 1 {
			routes := strings.Split(v.Mark, AuthUserRulesMarkSeparator)
			for i := range routes {
				var methods []string
				if method := v.Methods; method != "" {
					methods = []string{method}
				} else {
					methods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
				}
				setData(methods, routes[i], v)
			}
		} else {
			s.Marks = append(s.Marks, v.Mark)
		}
	}
	return s
}

func (*migrate) CreateAuthUserGroup() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthUserGroup", func(db *gorm.DB) error {
			data := []AuthUserGroup{
				{
					Name:   "管理员",
					Remark: "管理员",
					Status: 1,
					ID:     1,
				},
				{
					Name:   "编辑员",
					Remark: "编辑员",
					Status: 1,
					ID:     2,
				},
			}
			db.Create(data)
			return nil
		}
	})
}

// 获取权限标识列表
func (g AuthUserGroup) GetMarks() []string {
	currentRules := g.GetRules()
	res := make([]string, 0)
	for _, v := range currentRules {
		if TYPE_MARK == v.Type {
			res = append(res, v.Mark)
		}
	}

	return res;
}
