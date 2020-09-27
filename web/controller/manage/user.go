package manage

import (
	"github.com/sohaha/zlsgo/znet"
)

type User struct {}


func (*User) PostGetToken(c *znet.Context) {
	c.ApiJSON(200, "dd", nil)
}

func (*User) GetUseriInfo(c *znet.Context) {

}

func (*User) PutUpdate(c *znet.Context) {

}

func (*User) PutEditPassword(c *znet.Context) {

}

func (*User) PostUploadAvatar(c *znet.Context) {

}

func (*User) PostClearToken(c *znet.Context) {

}

func Create() {
}
