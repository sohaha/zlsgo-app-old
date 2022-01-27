package logic

import (
	"github.com/sohaha/zlsgo/zstring"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword(zstring.String2Bytes(hash), zstring.String2Bytes(password))
	return err == nil
}
