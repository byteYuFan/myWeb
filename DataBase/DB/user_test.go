package db

import (
	"context"
	"testing"
)

func TestGetUserInfoByUserName(t *testing.T) {
	Init()
	ctx := context.Background()
	res, err := GetUserInfoByUserName(ctx, "wyf")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}

func TestGetUserInfoByUserId(t *testing.T) {
	Init()
	ctx := context.Background()
	res, err := GetUserInfoByUserId(ctx, int64(1))
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}

func TestGetUserInfoByUserEmail(t *testing.T) {
	Init()
	ctx := context.Background()
	res, err := GetUserInfoByUserEmail(ctx, "854978151@qq.com")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}

func TestCreateUser(t *testing.T) {
	Init()
	user := &User{
		Username: "pogf",
		Password: "123456",
		Email:    "2541564974@qq.com",
	}
	err := CreateUser(context.Background(), user)
	if err != nil {
		t.Fatal(err)
	}

}

func TestMGetUsers(t *testing.T) {
	Init()
	ctx := context.Background()
	userIDs := []int64{
		1, 2,
	}
	res, err := MGetUsers(ctx, userIDs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res[0])
}
