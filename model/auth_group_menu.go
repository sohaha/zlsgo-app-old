package model

import (
	"app/web/business/manageBusiness"
	"errors"
	"gorm.io/gorm"
	"sort"
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
			return errors.New("服务繁忙,请重试")
		}
	} else {
		if res := db.Model(&m).Select("update_time", "menu").Where("id = ?", hasInfo.ID).Updates(m); res.RowsAffected == 0 {
			return errors.New("服务繁忙,请重试")
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

func (m *AuthGroupMenu) conv(menu Menu, user *AuthUser) (r Router) {
	show := manageBusiness.InArray(append(strings.Split(m.Menu, ","), "1"), strconv.Itoa(int(menu.ID)))
	has := manageBusiness.InArray(append(strings.Split(m.Menu, ","), "1", "2", "7"), strconv.Itoa(int(menu.ID)))
	if user.IsSuper {
		has = true
		show = true
	}

	r = Router{
		Name: menu.Title,
		Path: m.VuePath(menu.Index),
		// Url:        m.VueUrl(manageBusiness.InArray(strings.Split(m.Menu, ","), strconv.Itoa(int(menu.ID))), menu.Index),
		Url:      m.VueUrl(true, menu.Index),
		Icon:     menu.Icon,
		Children: []Router{},
	}
	r.Meta.Breadcrumb = menu.Breadcrumb == 1
	r.Meta.Real = menu.Real == 1
	r.Meta.Show = menu.Show == 1 && show
	r.Meta.Has = has
	r.Meta.Collapse = false

	return
}

func (m *AuthGroupMenu) getChild(menu Menu, menus []Menu, user *AuthUser) []Router {
	re := make([]Router, 0)
	for _, v := range menus {
		if v.Pid == uint8(menu.ID) {
			re = append(re, m.conv(v, user))
		}
	}

	return re
}

func (m *AuthGroupMenu) MenuInfo(user *AuthUser) (re []Router) {
	menuInfo := (&Menu{}).All()

	var groupIDArr []string
	for _, groupID := range user.GroupID {
		groupIDArr = append(groupIDArr, strconv.Itoa(int(groupID)))
	}

	var res []AuthGroupMenu
	menuKV := map[string]uint{}
	db.Where("groupid IN ?", groupIDArr).Find(&res)
	for _, groupRes := range res {
		for _, m := range strings.Split(groupRes.Menu, ",") {
			menuKV[m] = 1
		}
	}

	var mergeMenu []int
	for gid := range menuKV {
		if g, err := strconv.Atoi(gid); err == nil {
			mergeMenu = append(mergeMenu, g)
		}
	}

	var mergeMenuStr []string
	sort.Ints(mergeMenu)
	for _, gid := range mergeMenu {
		mergeMenuStr = append(mergeMenuStr, strconv.Itoa(gid))
	}
	m.Menu = strings.Join(mergeMenuStr, ",")

	for _, menu := range menuInfo {
		if menu.Pid == 0 {
			r := m.conv(menu, user)
			r.Children = m.getChild(menu, menuInfo, user)
			for _, mm := range r.Children {
				if mm.Meta.Show {
					r.Meta.Collapse = true
				}

				if mm.Meta.Has && r.Name != "后台中心" {
					r.Url = ""
				}
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

	if url == "main" {
		return "pages/main/" + url + ".vue"
	}

	if strings.HasPrefix(url, "/") {
		return "pages" + url + ".vue"
	}

	return "pages/" + url + ".vue"
}

func (m *AuthGroupMenu) VuePath(path string) string {
	if strings.HasPrefix(path, "/") {
		if !strings.HasPrefix(path, "/main") {
			return "/main" + path
		}
	} else {
		if path == "main" {
			return "/" + path + "/main"
		} else if !strings.HasPrefix(path, "/main/") {
			return "/main/" + path
		}
	}

	return path
}
