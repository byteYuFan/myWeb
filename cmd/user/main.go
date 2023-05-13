package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"myWeb/cmd/user/command"
	"myWeb/kitex_gen/user/usersrv/registerservice"
	"net"
)

var (
	Argon2Config *command.Argon2Params
)

func init() {
	Argon2Config = &command.Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
}
func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8001")
	s := registerservice.NewServer(new(MyRegisterServiceServer), server.WithServiceAddr(addr))
	if err := s.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", "登录", err)
	}
}
