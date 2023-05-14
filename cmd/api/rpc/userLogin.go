package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/kitex_gen/user/usersrv/loginservice"
	"myWeb/pkg/errno"
	"time"
)

var userLoginClient loginservice.Client

func Login(ctx context.Context, req *usersrv.UsernamePasswordLoginRequest) (resp *usersrv.UsernamePasswordLoginResponse, err error) {
	resp, err = userLoginClient.UsernamePasswordLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}
func SendEmail(ctx context.Context, req *usersrv.SendEmailRequest) (resp *usersrv.SendEmailResponse, err error) {
	resp, err = userLoginClient.SendEmail(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}
func LoginByEmail(ctx context.Context, req *usersrv.EmailLoginRequest) (resp *usersrv.EmailLoginResponse, err error) {
	resp, err = userLoginClient.EmailLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 10000 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}

func initLogin() {
	c, err := loginservice.NewClient("注册服务",
		client.WithHostPorts("127.0.0.1:8082"),
		client.WithRPCTimeout(30*time.Second),
		client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
			MaxIdlePerAddress: 1000,
			MaxIdleTimeout:    time.Minute}),
	)
	if err != nil {
		panic(err)
	}
	userLoginClient = c
}
func init() {
	initLogin()
}
