package handles

import (
	"github.com/gin-gonic/gin"
	"log"
	"myWeb/DataBase/pack"
	"myWeb/cmd/api/rpc"
	"myWeb/kitex_gen/userInfo"
	"myWeb/pkg/errno"
)

func RestPassword(ctx *gin.Context) {
	email := ctx.PostForm("Email")
	code := ctx.PostForm("Credential")
	resp, err := rpc.RestPassword(ctx, &userInfo.ResetPasswordRequest{
		Email:      email,
		Credential: code,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildRestPasswordResp(err))
		return
	}
	SendResponse(ctx, resp)
}

func UpdateUserInfo(ctx *gin.Context) {
	var userInformation = &userInfo.UpdateUserRequest{}
	i, _ := ctx.Get("id")
	id := i.(int64)
	userInformation.Id = id
	err := ctx.ShouldBind(userInformation)
	if err != nil {
		SendResponse(ctx, pack.BuildUpdateUserInfoResp(errno.ErrBind))
		return
	}
	resp, err := rpc.UpdateUserInfo(ctx, userInformation)
	if err != nil {
		SendResponse(ctx, pack.BuildUpdateUserInfoResp(err))
		return
	}
	SendResponse(ctx, resp)

}

func ChangeUserPassword(ctx *gin.Context) {
	var changePassword = new(userInfo.ChangePasswordRequest)
	err := ctx.ShouldBind(changePassword)
	if err != nil {
		SendResponse(ctx, pack.BuildChangeUserPasswordResp(errno.ErrBind))
		return
	}
	i, _ := ctx.Get("id")
	id := i.(int64)
	changePassword.Id = id
	log.Println(changePassword)
	resp, err := rpc.ChangeUserInfo(ctx, changePassword)
	if err != nil {
		SendResponse(ctx, pack.BuildChangeUserPasswordResp(err))
		return
	}
	SendResponse(ctx, resp)

}
