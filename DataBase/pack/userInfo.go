package pack

import (
	"github.com/pkg/errors"
	"myWeb/kitex_gen/userInfo"
	"myWeb/pkg/errno"
)

func BuildGetUserInfoResponse(err error) *userInfo.GetUserResponse {
	if err == nil {
		return getUserInfoResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getUserInfoResp(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return getUserInfoResp(s)
}
func getUserInfoResp(err errno.ErrNo) *userInfo.GetUserResponse {
	return &userInfo.GetUserResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
		User:        nil,
	}
}
func BuildRestPasswordResp(err error) *userInfo.ResetPasswordResponse {
	if err == nil {
		return restPasswordResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return restPasswordResp(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return restPasswordResp(s)
}
func restPasswordResp(err errno.ErrNo) *userInfo.ResetPasswordResponse {
	return &userInfo.ResetPasswordResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
	}
}
func updateUserInfoResp(err errno.ErrNo) *userInfo.UpdateUserResponse {
	return &userInfo.UpdateUserResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
	}
}
func BuildUpdateUserInfoResp(err error) *userInfo.UpdateUserResponse {
	if err == nil {
		return updateUserInfoResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return updateUserInfoResp(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return updateUserInfoResp(s)
}

func changeUserPasswordResp(err errno.ErrNo) *userInfo.ChangePasswordResponse {
	return &userInfo.ChangePasswordResponse{
		StatusCode:  int32(err.ErrCode),
		Description: err.ErrMsg,
	}
}
func BuildChangeUserPasswordResp(err error) *userInfo.ChangePasswordResponse {
	if err == nil {
		return changeUserPasswordResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return changeUserPasswordResp(e)
	}
	s := errno.ErrUnknown.WithMessage(err.Error())
	return changeUserPasswordResp(s)
}
