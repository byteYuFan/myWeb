package cmd

import (
	"context"
	"fmt"
	"myWeb/kitex_gen/userInfo"
	"testing"
)

func TestUserInfoInstance_ResetPassword(t *testing.T) {
	_, err := NewUserInfoInstance(context.Background()).ResetPassword(&userInfo.ResetPasswordRequest{
		Email:      "854978151@qq.com",
		Credential: "12867",
	}, &Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	})
	if err != nil {
		fmt.Println(err)
	}

}

func TestUserInfoInstance_ChangeUserPassword(t *testing.T) {
	err := NewUserInfoInstance(context.Background()).ChangeUserPassword(
		&userInfo.ChangePasswordRequest{
			Id:                 18,
			OldPassword:        "wwyyff123123456",
			NewPassword:        "wyf5211314",
			ConfirmNewPassword: "wyf5211314",
		}, &Argon2Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		})
	if err != nil {
		t.Log(err)
		return
	}
}
