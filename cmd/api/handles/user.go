package handles

import (
	"context"
	"fmt"
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
	fmt.Println("************", registerVar)
	resp, err := rpc.Register(context.Background(), &usersrv.RegisterRequest{
		Username:        registerVar.UserName,
		Email:           registerVar.Email,
		Password:        registerVar.Password,
		ConfirmPassword: registerVar.Confirm,
		Code:            registerVar.Code,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildUserRegisterResponse(err))
		return
	}
	SendResponse(ctx, resp)
}

func Login(ctx *gin.Context) {
	var loginVar UserLoginRequestParam
	loginVar.UserName = ctx.PostForm("UserName")
	loginVar.Password = ctx.PostForm("Password")
	resp, err := rpc.Login(ctx, &usersrv.UsernamePasswordLoginRequest{
		Username: loginVar.UserName,
		Password: loginVar.Password,
	})
	if err != nil {
		log.Println(err)
		SendResponse(ctx, pack.BuildUserRegisterResponse(err))
		return
	}
	SendResponse(ctx, resp)
}

func SendEmail(ctx *gin.Context) {
	email := ctx.PostForm("Email")
	resp, err := rpc.SendEmail(ctx, &usersrv.SendEmailRequest{
		Email: email,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildSendEmailResponse(err))
		return
	}
	SendResponse(ctx, resp)
}

func LoginByEmail(ctx *gin.Context) {
	email := ctx.PostForm("Email")
	code := ctx.PostForm("Credential")
	resp, err := rpc.LoginByEmail(ctx, &usersrv.EmailLoginRequest{
		Email:      email,
		Credential: code,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildEmailLoginResponse(err))
		return
	}
	SendResponse(ctx, resp)
}
