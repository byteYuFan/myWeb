package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/kitex_gen/user/usersrv/registerservice"
)

func main() {
	var opts = client.WithHostPorts("127.0.0.1:8001")
	newClient, _ := registerservice.NewClient("登录", opts)

	// Send a Register request
	req := &usersrv.RegisterRequest{
		Username:        "wyf",
		Email:           "testuser@example.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	resp, err := newClient.Register(context.Background(), req)
	if err != nil {
		panic(fmt.Sprintf("Failed to register user: %v", err))
	}

	fmt.Println(resp)
	fmt.Printf("User registered successfully with status code %d and description: %s", resp.StatusCode, resp.Description)
}
