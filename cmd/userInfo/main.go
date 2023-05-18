package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"myWeb/cmd/userInfo/cmd"
	"myWeb/kitex_gen/userInfo/userservice"
	"myWeb/pkg/ttviper"
	"net"
)

var (
	cfg          = ttviper.ConfigInit("userInfo.yml")
	serverHost   = cfg.Viper.GetString("UserInfoServer.Addr")
	serverPort   = cfg.Viper.GetString("UserInfoServer.Port")
	Argon2Config *cmd.Argon2Params
)

func init() {
	Argon2Config = &cmd.Argon2Params{
		Memory:      cfg.Viper.GetUint32("UserInfoServer.Argon2ID.Memory"),
		Iterations:  cfg.Viper.GetUint32("UserInfoServer.Argon2ID.Iterations"),
		Parallelism: uint8(cfg.Viper.GetUint("UserInfoServer.Argon2ID.Parallelism")),
		SaltLength:  cfg.Viper.GetUint32("UserInfoServer.Argon2ID.SaltLength"),
		KeyLength:   cfg.Viper.GetUint32("UserInfoServer.Argon2ID.KeyLength"),
	}
}
func main() {
	addr, _ := net.ResolveTCPAddr("tcp", net.JoinHostPort(serverHost, serverPort))
	s := userservice.NewServer(new(UserInfoServiceServer), server.WithServiceAddr(addr))
	if err := s.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", "user-info", err)
	}
}
