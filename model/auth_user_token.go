package model

import (
	"gorm.io/gorm"
)

// AuthUserToken 管理员密钥
type AuthUserToken struct {
	gorm.Model
	Userid int
	Token  string
	IP     string
	UA     string
	Status uint8
}
