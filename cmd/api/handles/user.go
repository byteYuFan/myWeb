package handles

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"myWeb/DataBase/pack"
	"myWeb/cmd/api/rpc"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/pkg/errno"
)

func Register(ctx *gin.Context) {
	var registerVar UserRegisterRequestParam
	err := ctx.ShouldBind(&registerVar)
	if err != nil {
		SendResponse(ctx, pack.BuildUserRegisterResponse(errno.ErrBind))
		return
	}
	log.Println(registerVar)
	resp, err := rpc.Register(context.Background(), &usersrv.RegisterRequest{
		Username:        registerVar.UserName,
		Email:           registerVar.Email,
		Password:        registerVar.Password,
		ConfirmPassword: registerVar.Confirm,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildUserRegisterResponse(err))
		return
	}
	SendResponse(ctx, resp)
}
