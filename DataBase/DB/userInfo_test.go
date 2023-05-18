package db

import (
	"context"
	"testing"
)

func TestUpdateUserInfo(t *testing.T) {
	user := &UserInfo{
		Id:         0,
		Name:       "wyf",
		Password:   "",
		Age:        0,
		Profession: "",
		Department: "",
		Province:   "山西",
		City:       "",
		Flag:       true,
	}
	err := UpdateUserInfo(context.Background(), 16, user)
	if err != nil {
		return
	}
}
