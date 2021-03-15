package model

import (
	"errors"
	"gorm.io/gorm"
)

// AuthGroupMenu 角色菜单对应表
type AuthGroupMenu struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	GroupID   uint8          `gorm:"column:groupid;type:tinyint(4);not null;default:0;comment:角色Id;" json:"groupid"`
	Menu      string         `gorm:"column:menu;type:varchar(255);not null;default:'';comment:菜单;" json:"menu"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

func (*migrate) CreateAuthGroupMenu() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthGroupMenu", func(db *gorm.DB) error {
			tx := db.Create([]AuthGroupMenu{
				{
					GroupID: 1,
					Menu:    "1,2,3,4,5,6,7,8,9,10,11,12",
				},
				{
					GroupID: 2,
					Menu:    "1",
				},
			})
			return tx.Error
		}
	})
}

func (m *AuthGroupMenu) Update() error {
	hasInfo := &AuthGroupMenu{}
	db.Where("groupid = ?", m.GroupID).First(hasInfo)
	if hasInfo.ID == 0 {
		if res := db.Create(m); res.RowsAffected == 0 {
			return errors.New("服务繁忙，请重试")
		}
	} else {
		if res := db.Model(&m).Select("update_time", "menu").Where("id = ?", hasInfo.ID).Updates(m); res.RowsAffected == 0 {
			return errors.New("服务繁忙，请重试")
		}
	}

	return nil
}

type Router struct {
	ID   uint   `json:"-"`
	Name string `json:"name"`
	Path string `json:"path"`
	Url  string `json:"url"`
	Icon string `json:"icon"`
	Meta struct {
		Breadcrumb bool `json:"breadcrumb"`
		Real       bool `json:"real"`
		Show       bool `json:"show"`
		Has        bool `json:"has"`
		Collapse   bool `json:"collapse"`
	} `json:"meta"`
	Children []Router `json:"children"`
}

func (m *AuthGroupMenu) SelectGroupMenu(groupid uint8) AuthGroupMenu {
	groupMenuInfo := AuthGroupMenu{}
	db.Model(&AuthGroupMenu{}).Where("groupid = ?", groupid).Find(&groupMenuInfo)

	return groupMenuInfo
}
