package model

import (
	"gorm.io/gorm"
)

// Setting 系统设置表
type Setting struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Varname   string         `gorm:"column:varname;unique;type:varchar(255);not null;default:'';comment:配置名称;" json:"varname"`
	Value     string         `gorm:"column:value;type:varchar(255);not null;default:'';comment:配置值;" json:"value"`
	Info      string         `gorm:"column:info;type:varchar(255);not null;default:'';comment:配置说明;" json:"info"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

func (*migrate) CreateSetting() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateSetting", func(db *gorm.DB) error {
			initData := []Setting{
				{
					Varname: "sitename",
					Value:   "内容管理系统",
					Info:    "网站标题",
				},
				{
					Varname: "domain",
					Value:   "",
					Info:    "网站域名",
				},
				{
					Varname: "login_expire_time",
					Value:   "3600",
					Info:    "登录超时时间",
				},
				{
					Varname: "login_mode",
					Value:   "0",
					Info:    "登录模式",
				},
			}
			tx := db.Create(initData)
			if tx.Error != nil {
				log.Fatalf("初始化系统设置失败: %s", tx.Error.Error())
			}
			return nil
		}
	})
}

func (s *Setting) SettingValues() []Setting {
	var res []Setting
	db.Find(&res)

	return res
}

func (s *Setting) SettingValue() {
	db.Where(s, "varname").Find(s)
}

func (s *Setting) SettingInsert() error {
	return db.Create(s).Error
}

func (s *Setting) SettingUpdate() error {
	return db.Save(s).Error
}
