package model

import (
	"errors"
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
	return db.Where(&AuthUserGroup{Status: g.Status}).Order("id desc").Find(&groups)
}

var ruleCache = zcache.New("ruleCache")

type GetRulesModel struct {
	AuthUserRules
	RelaStauts uint8 `json:"rela_stauts"`
}

// GetRules 获取规则列表
func (g AuthUserGroup) GetRules() (rules []GetRulesModel) {
	currentRule, err := ruleCache.MustGet(strconv.Itoa(int(g.ID)), func(set func(data interface{}, lifeSpan time.Duration, interval ...bool)) (err error) {
		var relas []AuthUserRulesRela
		// db.Model(&AuthUserRulesRela{}).Where(map[string]interface{}{"group_id": g.ID, "status": 1}).Select("rule_id").Find(&relas)
		db.Model(&AuthUserRulesRela{}).Where(map[string]interface{}{"group_id": g.ID}).Select("rule_id").Find(&relas)
		var ids []uint

		for _, v := range relas {
			ids = append(ids, v.RuleID)
		}

		authUserRulesTable := TableName("auth_user_rules")
		authUserRulesRelaTable := TableName("auth_user_rules_rela")

		db.Model(&AuthUserRules{}).Select(authUserRulesTable+".*", authUserRulesRelaTable+".status as rela_stauts").Where(authUserRulesRelaTable+".group_id = ? and "+authUserRulesTable+".id in (?)", g.ID, ids).Joins("LEFT JOIN " + authUserRulesRelaTable + " ON " + authUserRulesRelaTable + ".rule_id = " + authUserRulesTable + ".id").Scan(&rules)

		set(rules, 60*10, true)
		return nil
	})
	if err != nil {
		return
	}
	return currentRule.([]GetRulesModel)
}

// GetRuleCollation 获取整理后的规则列表
func (g AuthUserGroup) GetRuleCollation(p *RuleCollation) (s *RuleCollation) {
	rules := g.GetRules()
	if len(rules) == 0 {
		return
	}

	if p != nil {
		s = p
	} else {
		// 有必要可以升级成树
		s = &RuleCollation{
			AdoptRoute:     map[string][]string{},
			InterceptRoute: map[string][]string{},
			Marks:          []string{},
		}
	}

	setData := func(methods []string, route string, v GetRulesModel) {
		for _, m := range methods {
			if v.RelaStauts == AuthUserRulesStatusIntercept {
				s.InterceptRoute[m] = append(s.InterceptRoute[m], route)
			} else if v.RelaStauts == AuthUserRulesStatusAdopt {
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
			if v.Status == 1 {
				s.Marks = append(s.Marks, v.Mark)
			}
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
		if TYPE_MARK == v.Type && AuthUserRulesStatusAdopt == v.RelaStauts {
			res = append(res, v.Mark)
		}
	}

	return res
}

func (g *AuthUserGroup) GroupInfo() {
	_ = db.First(&g, g.ID)
}

func (g *AuthUserGroup) Exist() error {
	var res *gorm.DB
	if g.ID == 0 {
		res = db.Where("name = ?", g.Name).Find(&AuthUserGroup{})
	} else {
		res = db.Where("name = ? and id != ?", g.Name, g.ID).Find(&AuthUserGroup{})
	}
	if res.RowsAffected == 0 {
		return nil
	}

	return errors.New("角色名称已存在")
}

func (g *AuthUserGroup) Save() error {
	if g.ID == 0 {
		if res := db.Create(&g); res.RowsAffected == 0 {
			return errors.New("保存失败")
		}
	} else {
		res := db.Model(&g).Select([]string{"name", "remark", "update_time"}).Updates(g)
		if res.RowsAffected == 0 {
			return errors.New("保存失败")
		}
	}

	return nil
}
