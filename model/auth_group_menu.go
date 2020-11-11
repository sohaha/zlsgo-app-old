package model

import (
	"app/web/business/manageBusiness"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
			db.Create([]AuthGroupMenu{
				{
					GroupID: 1,
					Menu:    "1,2,3,4,5,6,7,8,9,10,11",
				},
				{
					GroupID: 2,
					Menu:    "1",
				},
			})
			return nil
		}
	})
}

func (m *AuthGroupMenu) Update() error {
	hasInfo := &AuthGroupMenu{}
	db.Where("groupid = ?", m.GroupID).First(hasInfo)
	if hasInfo.ID == 0 {
		if res := db.Create(m); res.RowsAffected == 0 {
			return errors.New("服务繁忙,请重试.")
		}
	} else {
		if res := db.Model(&m).Select("update_time", "menu").Where("id = ?", hasInfo.ID).Updates(m); res.RowsAffected == 0 {
			return errors.New("服务繁忙,请重试.")
		}
	}

	return nil
}

type Router struct {
	ID         uint     `json:"-"`
	Name       string   `json:"name"`
	Path       string   `json:"path"`
	Url        string   `json:"url"`
	Icon       string   `json:"icon"`
	Breadcrumb bool     `json:"breadcrumb"`
	Real       bool     `json:"real"`
	Show       bool     `json:"show"`
	Has        bool     `json:"has"`
	Collapse   bool     `json:"collapse"`
	Children   []Router `json:"children"`
}

func (m *AuthGroupMenu) conv(menu Menu) Router {
	return Router{
		Name: menu.Title,
		Path: m.VuePath(menu.Index),
		// Url:        m.VueUrl(manageBusiness.InArray(strings.Split(m.Menu, ","), strconv.Itoa(int(menu.ID))), menu.Index),
		Url:        m.VueUrl(true, menu.Index),
		Icon:       menu.Icon,
		Breadcrumb: menu.Breadcrumb == 1,
		Real:       menu.Real == 1,
		Show:       menu.Show == 1,
		Has:        manageBusiness.InArray(append(strings.Split(m.Menu, ","), "1", "2", "3", "4"), strconv.Itoa(int(menu.ID))),
		Collapse:   false,
		Children:   []Router{},
	}
}

func (m *AuthGroupMenu) getChild(menu Menu, menus []Menu) []Router {
	re := make([]Router, 0)
	for _, v := range menus {
		if v.Pid == uint8(menu.ID) {
			re = append(re, m.conv(v))
		}
	}

	return re
}

func (m *AuthGroupMenu) MenuInfo() (re []Router) {
	menuInfo := (&Menu{}).All()
	db.Where("groupid = ?", m.GroupID).Find(&m)

	for _, menu := range menuInfo {
		if menu.Pid == 0 {
			r := m.conv(menu)
			r.Children = m.getChild(menu, menuInfo)
			for _, mm := range r.Children {
				if mm.Show {
					r.Collapse = true
				}
			}
			if r.Collapse {
				r.Url = ""
			}

			re = append(re, r)
		}
	}

	return re
}

func (m *AuthGroupMenu) VueUrl(show bool, url string) string {
	if !show {
		return ""
	}
	return "pages/main/" + url + ".vue"
}

func (m *AuthGroupMenu) VuePath(path string) string {
	if strings.HasPrefix(path, "/") {
		if !strings.HasPrefix(path, "/main") {
			return "/main" + path
		}
	} else {
		if path == "main" {
			return "/" + path
		} else if !strings.HasPrefix(path, "/main/") {
			return "/main/" + path
		}
	}

	return path
}
