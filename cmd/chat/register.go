package main

import (
	"context"
	"myWeb/kitex_gen/chatgpt"
)

type MyChatServiceServer struct {
}

func (chat *MyChatServiceServer) Chat(ctx context.Context, req *chatgpt.ChatRequest) (res *chatgpt.ChatResponse, err error) {
	return nil, nil
}
