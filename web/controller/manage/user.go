package manage

import (
	"context"

	"app/global"
	"app/global/ent"
	"app/model/authuser"
	"github.com/sohaha/zlsgo/znet"
)

type User struct {
}

func (*User) PostGetToken(c *znet.Context) {
	client := ent.Client()

	client.AuthUser.Query().Where(authuser.Username(""))
	Create()
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
	client := ent.Client()
	u, err := client.AuthUser.Create().
		SetUsername("hi").
		SetPassword("33").
		Save(context.Background())
	global.Log.Debug(u, err)
}
