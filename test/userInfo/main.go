package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"myWeb/kitex_gen/userInfo"
	"myWeb/kitex_gen/userInfo/userservice"
)

func main() {
	var opts = client.WithHostPorts("127.0.0.1:8083")
	newClient, _ := userservice.NewClient("获取用户信息", opts)
	req := &userInfo.GetUserRequest{
		Id: 16,
	}
	resp, err := newClient.GetUser(context.Background(), req)
	if err != nil {
		panic(fmt.Sprintf("Failed to chat: %v", err))
	}

	fmt.Println(resp)
	fmt.Println(resp.User)

}
