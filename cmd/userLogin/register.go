package main

import (
	"context"
	"myWeb/DataBase/pack"
	"myWeb/cmd/userLogin/command"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/pkg/checkout"
	"myWeb/pkg/errno"
	"myWeb/pkg/jwt"
)

type MyLoginServiceServer struct {
}

func (login *MyLoginServiceServer) UsernamePasswordLogin(ctx context.Context, req *usersrv.UsernamePasswordLoginRequest) (resp *usersrv.UsernamePasswordLoginResponse, err error) {
	if !checkout.ValidateUsername(req.Username) {
		resp = pack.BuildUsernameLoginResponse(errno.ErrInvalidUsername)
		return
	}
	if !checkout.ValidatePassword(req.Password) {
		resp = pack.BuildUsernameLoginResponse(errno.ErrInvalidPassword)
		return
	}
	uid, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuildUsernameLoginResponse(err)
		return resp, nil
	}
	token, err := jwt.GenerateToken(uid, req.Username)
	if err != nil {
		resp = pack.BuildUsernameLoginResponse(errno.ErrSignatureInvalid)
		return resp, nil
	}
	resp = &usersrv.UsernamePasswordLoginResponse{
		StatusCode:  10000,
		Description: "login successfully!",
		Token:       token,
	}
	return resp, nil
}

func (login *MyLoginServiceServer) EmailLogin(ctx context.Context, req *usersrv.EmailLoginRequest) (resp *usersrv.EmailLoginResponse, err error) {
	if !checkout.ValidateEmail(req.Email) {
		resp = pack.BuildEmailLoginResponse(errno.ErrInvalidEmail)
		return
	}
	uid, username, err := command.NewCheckUserService(ctx).CheckUserEmail(req)
	if err != nil {
		resp = pack.BuildEmailLoginResponse(err)
		return
	}
	token, err := jwt.GenerateToken(uid, username)
	if err != nil {
		resp = pack.BuildEmailLoginResponse(errno.ErrSignatureInvalid)
		return resp, nil
	}
	resp = &usersrv.EmailLoginResponse{
		StatusCode:  10000,
		Description: "login successfully!",
		Token:       token,
	}
	return resp, nil
}

func (login *MyLoginServiceServer) SendEmail(ctx context.Context, req *usersrv.SendEmailRequest) (resp *usersrv.SendEmailResponse, err error) {
	if !checkout.ValidateEmail(req.Email) {
		resp = pack.BuildSendEmailResponse(errno.ErrInvalidEmail)
		return
	}
	_, err = command.NewSendEmailService(ctx).Send(req.Email)
	if err != nil {
		resp = pack.BuildSendEmailResponse(err)
		return
	}
	resp = pack.BuildSendEmailResponse(errno.Success)
	resp.StatusCode = 10000
	resp.Description = "send email successfully!"
	return resp, nil
}
