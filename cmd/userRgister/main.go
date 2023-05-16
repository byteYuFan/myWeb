package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"myWeb/cmd/userRgister/command"
	"myWeb/kitex_gen/user/usersrv/registerservice"
	"myWeb/pkg/ttviper"
	"net"
)

var (
	cfg           = ttviper.ConfigInit("userConfig.yml")
	ServerAddress = cfg.Viper.GetString("RegisterServer.Address")
	ServerPort    = cfg.Viper.GetString("RegisterServer.Port")
	Argon2Config  *command.Argon2Params
)

func init() {
	Argon2Config = &command.Argon2Params{
		Memory:      cfg.Viper.GetUint32("RegisterServer.Argon2ID.Memory"),
		Iterations:  cfg.Viper.GetUint32("RegisterServer.Argon2ID.Iterations"),
		Parallelism: uint8(cfg.Viper.GetUint("RegisterServer.Argon2ID.Parallelism")),
		SaltLength:  cfg.Viper.GetUint32("RegisterServer.Argon2ID.SaltLength"),
		KeyLength:   cfg.Viper.GetUint32("RegisterServer.Argon2ID.KeyLength"),
	}
}
func main() {
	addr, _ := net.ResolveTCPAddr("tcp", net.JoinHostPort(ServerAddress, ServerPort))
	s := registerservice.NewServer(new(MyRegisterServiceServer), server.WithServiceAddr(addr))
	if err := s.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", "登录", err)
	}
}
