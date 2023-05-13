package main

import (
	"context"
	"errors"
	"log"
	"myWeb/DataBase/pack"
	"myWeb/cmd/user/command"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/pkg/checkout"
	"myWeb/pkg/errno"
)

type MyRegisterServiceServer struct{}

func (s *MyRegisterServiceServer) Register(ctx context.Context, req *usersrv.RegisterRequest) (resp *usersrv.RegisterResponse, err error) {
	log.Println("-----------------", req)
	if checkout.ValidateUsername(req.Username) == false {
		resp = pack.BuildUserRegisterResponse(errno.ErrInvalidUsername)
		err = errors.New("username invalid")
		return
	}
	if checkout.ValidatePassword(req.Password) == false {
		resp = pack.BuildUserRegisterResponse(errno.ErrInvalidPassword)
		err = errors.New("password invalid")
		return
	}
	if checkout.ValidateEmail(req.Email) == false {
		resp = pack.BuildUserRegisterResponse(errno.ErrInvalidEmail)
		err = errors.New("email invalid")
		return
	}
	if req.Password != req.ConfirmPassword {
		resp = pack.BuildUserRegisterResponse(errors.New("the password is not same as confirm password"))
		err = errors.New("confirm password invalid")
		return
	}
	err = command.NewCreateUserService(ctx).CreateUser(req, Argon2Config)
	if err != nil {
		resp = pack.BuildUserRegisterResponse(err)
		return resp, nil
	}
	resp = pack.BuildUserRegisterResponse(errno.Success)
	resp.StatusCode = 0
	resp.Description = "register successfully."
	return resp, nil
}

func (s *MyRegisterServiceServer) MustEmbedUnimplementedRegisterServiceServer() {}
