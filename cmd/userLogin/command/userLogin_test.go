package command

import (
	"context"
	"fmt"
	"myWeb/kitex_gen/user/usersrv"
	"testing"
)

func TestCheckUserService_CheckUserEmail(t *testing.T) {
	req := &usersrv.EmailLoginRequest{
		Email:      "w1213ph@gmail.com",
		Credential: "666666",
	}
	i, j, k := NewCheckUserService(context.Background()).CheckUserEmail(req)
	fmt.Println(i, j, k)
}
func TestSend(t *testing.T) {
	NewSendEmailService(context.Background()).Send("850021638@qq.com")
}
