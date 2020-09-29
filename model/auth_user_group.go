package model

import (
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zcache"
)

type AuthUserGroup struct {
	Name      string         `gorm:"type:varbinary(100);default:''"`
	Remark    string         `gorm:"type:varbinary(250);default:''"`
	Status    uint8          `gorm:"type:int(2);default:1"`
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt JSONTime       `gorm:"column:create_time;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
			data := []AuthUserGroup{
				{
					Name:   "管理员",
					Remark: "我是一个管理员",
					Status: 1,
					ID:     1,
				},
				{
					Name:   "编辑员",
					Remark: "我是一个编辑员",
					Status: 1,
					ID:     2,
				},
			}
			db.Create(data)
			return nil
		}
	})
}
