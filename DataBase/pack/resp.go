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
func BuildUsernameLoginResponse(err error) *usersrv.UsernamePasswordLoginResponse {
	if err == nil {
		return userLoginResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userLoginResp(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return userLoginResp(s)

}
func userLoginResp(err errno.ErrNo) *usersrv.UsernamePasswordLoginResponse {
	return &usersrv.UsernamePasswordLoginResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
		Token:       "",
	}
}

func BuildSendEmailResponse(err error) *usersrv.SendEmailResponse {
	if err == nil {
		return emailResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return emailResp(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return emailResp(s)
}
func emailResp(err errno.ErrNo) *usersrv.SendEmailResponse {
	return &usersrv.SendEmailResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
	}
}
func BuildEmailLoginResponse(err error) *usersrv.EmailLoginResponse {
	if err == nil {
		return emailLogin(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return emailLogin(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return emailLogin(s)

}
func emailLogin(err errno.ErrNo) *usersrv.EmailLoginResponse {
	return &usersrv.EmailLoginResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
		Token:       "",
	}
}
