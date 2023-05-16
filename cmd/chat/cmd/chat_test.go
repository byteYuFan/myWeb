package cmd

import (
	"context"
	"myWeb/kitex_gen/chatgpt"
	"testing"
)

func TestChatToChatGPTService_ChatRequest(t *testing.T) {
	NEWChatToCHATGPTService(context.Background()).ChatRequest(&chatgpt.ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []*chatgpt.ChatMessage{
			{
				Role:    "user",
				Content: "你好",
			},
		},
	})
}
