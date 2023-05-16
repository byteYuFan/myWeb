package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"myWeb/kitex_gen/user/usersrv/loginservice"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", "8082"))
	s := loginservice.NewServer(new(MyLoginServiceServer), server.WithServiceAddr(addr))
	if err := s.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", "登录---", err)
	}
}
