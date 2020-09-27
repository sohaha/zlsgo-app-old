package model

import (
	"strconv"
	"time"

	"github.com/sohaha/zlsgo/zcache"
	"gorm.io/gorm"
)

type AuthUserGroup struct {
	Name   string `gorm:"type:varbinary(100);default:''"`
	Remark string `gorm:"type:varbinary(250);default:''"`
	Status uint8  `gorm:"type:int(2);default:1"`
	// Rules  []AuthUserRules `gorm111:"many2many:auto_group_rule;"`
	gorm.Model
}

// All 获取全部记录
func (g AuthUserGroup) All(groups *[]AuthUserGroup) *gorm.DB {
	return db.Where(&AuthUserGroup{Status: g.Status}).Find(&groups)
}

var ruleCache = zcache.New("ruleCache")

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

func (*migrate) CreateAuthUserGroup() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthUserGroup", func(db *gorm.DB) error {
			// 添加一个 id 为 0 的用户为游客
			db.Create(&AuthUserGroup{
				Name:   "游客",
				Remark: "我是一个游客",
				Status: 1,
			})
			db.Model(&AuthUserGroup{}).Where("id = ?", 1).Update("id", 0)

			data := []AuthUserGroup{
				{
					Name:   "管理员",
					Remark: "我是一个管理员",
					Status: 1,
					Model: gorm.Model{
						ID: 1,
					},
				},
				{
					Name:   "编辑员",
					Remark: "我是一个编辑员",
					Status: 1,
					Model: gorm.Model{
						ID: 2,
					},
				},
			}
			db.Create(data)
			return nil
		}
	})
}
