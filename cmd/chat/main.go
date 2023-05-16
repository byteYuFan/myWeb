package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"myWeb/kitex_gen/chatgpt/chatservice"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", "8083"))
	s := chatservice.NewServer(new(MyChatServiceServer), server.WithServiceAddr(addr))
	if err := s.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", "chat-GPT", err)
	}
}
