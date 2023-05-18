package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"myWeb/kitex_gen/userInfo"
	"myWeb/kitex_gen/userInfo/userservice"
	"myWeb/pkg/errno"
	"time"
)

var userInfoClient userservice.Client

func RestPassword(ctx context.Context, req *userInfo.ResetPasswordRequest) (resp *userInfo.ResetPasswordResponse, err error) {
	resp, err = userInfoClient.RestPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}
func UpdateUserInfo(ctx context.Context, req *userInfo.UpdateUserRequest) (resp *userInfo.UpdateUserResponse, err error) {
	resp, err = userInfoClient.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}
func ChangeUserInfo(ctx context.Context, req *userInfo.ChangePasswordRequest) (resp *userInfo.ChangePasswordResponse, err error) {
	resp, err = userInfoClient.ChangePassword(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}
func initUserInfo() {
	c, err := userservice.NewClient("用户信息服务",
		client.WithHostPorts("127.0.0.1:8083"),
		client.WithRPCTimeout(30*time.Second),
		client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
			MaxIdlePerAddress: 1000,
			MaxIdleTimeout:    time.Minute}),
	)
	if err != nil {
		panic(err)
	}
	userInfoClient = c
}
func init() {
	initUserInfo()
}
