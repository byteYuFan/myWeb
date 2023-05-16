package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/kitex_gen/user/usersrv/loginservice"
)

func main() {
	var opts = client.WithHostPorts("127.0.0.1:8082")
	newClient, _ := loginservice.NewClient("登录", opts)

	// Send a Register request
	req := &usersrv.UsernamePasswordLoginRequest{
		Username: "wyfg4tgtt",
		Password: "password",
	}

	resp, err := newClient.UsernamePasswordLogin(context.Background(), req)
	if err != nil {
		panic(fmt.Sprintf("Failed to login: %v", err))
	}

	fmt.Println(resp)
	fmt.Printf("User registered successfully with status code %d and description: %s", resp.StatusCode, resp.Description)
}
