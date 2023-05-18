package main

import (
	"context"
	"errors"
	"math"
	"myWeb/DataBase/pack"
	"myWeb/cmd/userInfo/cmd"
	"myWeb/kitex_gen/userInfo"
	"myWeb/pkg/checkout"
	"myWeb/pkg/errno"
)

type UserInfoServiceServer struct {
}

// NewUserInfoServiceServer 返回一个UserInfoServiceServer的实例

func (info *UserInfoServiceServer) UpdateUser(ctx context.Context, req *userInfo.UpdateUserRequest) (resp *userInfo.UpdateUserResponse, err error) {
	if req.Id > math.MaxInt64 || req.Id <= 0 {
		resp = pack.BuildUpdateUserInfoResp(errno.NewErrNo(100701, "id不在允许的范围内"))
		return
	}
	err = cmd.NewUserInfoInstance(ctx).UpdateUserInfo(req)
	if err != nil {
		resp = pack.BuildUpdateUserInfoResp(err)
	}
	resp = pack.BuildUpdateUserInfoResp(errno.Success)
	resp.StatusCode = 10000
	resp.Description = "修改成功"
	return resp, nil
}

func (info *UserInfoServiceServer) ChangePassword(ctx context.Context, req *userInfo.ChangePasswordRequest) (resp *userInfo.ChangePasswordResponse, err error) {
	if !checkout.ValidatePassword(req.OldPassword) || !checkout.ValidatePassword(req.NewPassword) {
		resp = pack.BuildChangeUserPasswordResp(errno.ErrInvalidPassword)
		return
	}
	if req.NewPassword != req.ConfirmNewPassword {
		resp = pack.BuildChangeUserPasswordResp(errors.New("the password is not same as confirm password"))
		return
	}
	err = cmd.NewUserInfoInstance(ctx).ChangeUserPassword(req, Argon2Config)
	if err != nil {
		resp = pack.BuildChangeUserPasswordResp(err)
		return
	}
	resp = pack.BuildChangeUserPasswordResp(errno.Success)
	resp.StatusCode = 10000
	resp.Description = "修改成功"
	return resp, nil
}

func (info *UserInfoServiceServer) GetUser(ctx context.Context, req *userInfo.GetUserRequest) (resp *userInfo.GetUserResponse, err error) {
	if req.Id > math.MaxInt64 || req.Id <= 0 {
		resp = pack.BuildGetUserInfoResponse(errno.NewErrNo(100701, "id不在允许的范围内"))
		return
	}
	user, err := cmd.NewUserInfoInstance(ctx).GetUserInfo(req)
	if err != nil {
		resp = pack.BuildGetUserInfoResponse(err)
		return
	}
	resp = pack.BuildGetUserInfoResponse(errno.Success)
	resp.User = user
	return
}

func (info *UserInfoServiceServer) RestPassword(ctx context.Context, req *userInfo.ResetPasswordRequest) (resp *userInfo.ResetPasswordResponse, err error) {
	if !checkout.ValidateEmail(req.Email) {
		resp = pack.BuildRestPasswordResp(errno.ErrInvalidEmail)
		return
	}
	if len(req.Credential) != 6 {
		resp = pack.BuildRestPasswordResp(errno.NewErrNo(100602, "验证码位6位数"))
		return
	}
	_, err = cmd.NewUserInfoInstance(ctx).ResetPassword(req, Argon2Config)
	if err != nil {
		resp = pack.BuildRestPasswordResp(err)
		return
	}
	resp = pack.BuildRestPasswordResp(errno.Success)
	resp.StatusCode = 10000
	resp.Description = "修改成功默认密码为用户名+123456"
	return resp, nil
}
