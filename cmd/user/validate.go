package main

import (
	"github.com/pkg/errors"
	"myWeb/DataBase/pack"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/pkg/checkout"
	"myWeb/pkg/errno"
)

func validate(req *usersrv.RegisterRequest) (resp *usersrv.RegisterResponse, flag bool) {
	if checkout.ValidateUsername(req.Username) {
		resp = pack.BuildUserRegisterResponse(errno.ErrInvalidUsername)
		return
	}
	if checkout.ValidateEmail(req.Password) {
		resp = pack.BuildUserRegisterResponse(errno.ErrInvalidPassword)
		return
	}
	if checkout.ValidateEmail(req.Email) {
		resp = pack.BuildUserRegisterResponse(errno.ErrInvalidEmail)
		return
	}
	if req.Password != req.ConfirmPassword {
		resp = pack.BuildUserRegisterResponse(errors.New("the password is not same as confirm password"))
		return
	}
	flag = false
	return
}
