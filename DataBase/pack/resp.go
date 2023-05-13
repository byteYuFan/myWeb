package pack

import (
	"errors"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/pkg/errno"
)

func BuildUserRegisterResponse(err error) *usersrv.RegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return userRegisterResp(s)
}

func userRegisterResp(err errno.ErrNo) *usersrv.RegisterResponse {
	return &usersrv.RegisterResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
	}
}
