package db

import (
	"testing"
)

func TestInit(t *testing.T) {
	user := &User{
		Username: "wyf",
		Password: "123456",
		Email:    "854978151@qq.com",
	}
	USERDB.Create(user)
}
