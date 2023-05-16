package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"myWeb/kitex_gen/chatgpt"
	"myWeb/kitex_gen/chatgpt/chatservice"
)

func main() {
	var opts = client.WithHostPorts("127.0.0.1:8083")
	newclient, _ := chatservice.NewClient("chat", opts)
	req := &chatgpt.ChatRequest{
		Model:    "11111",
		Messages: nil,
	}
	resp, err := newclient.Chat(context.Background(), req)
	if err != nil {
		panic(fmt.Sprintf("Failed to chat: %v", err))
	}

	fmt.Println(resp)
	fmt.Printf("User registered successfully with status code %s and description: %s", resp.Id, resp.Object)

}
