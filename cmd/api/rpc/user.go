package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/kitex_gen/user/usersrv/registerservice"
	"myWeb/pkg/errno"
	"time"
)

var userRegisterClient registerservice.Client

func init() {
	c, err := registerservice.NewClient("注册服务",
		client.WithHostPorts("127.0.0.1:8001"),
		client.WithRPCTimeout(30*time.Second),
		client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
			MaxIdlePerAddress: 1000,
			MaxIdleTimeout:    time.Minute}),
	)
	if err != nil {
		panic(err)
	}
	userRegisterClient = c
}

func Register(ctx context.Context, req *usersrv.RegisterRequest) (resp *usersrv.RegisterResponse, err error) {
	resp, err = userRegisterClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), resp.Description)
	}
	return resp, nil
}
