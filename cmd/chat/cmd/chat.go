package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"myWeb/kitex_gen/chatgpt"
	"net/http"
)

type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index        int     `json:"index"`
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatToChatGPTService struct {
	ctx context.Context
}

func NEWChatToCHATGPTService(ctx context.Context) *ChatToChatGPTService {
	return &ChatToChatGPTService{
		ctx: ctx,
	}
}

func (chat *ChatToChatGPTService) ChatRequest(request *chatgpt.ChatRequest) (resp *chatgpt.ChatRequest, err error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 设置请求头中的Token认证
	req.Header.Set("Authorization", "Bearer  Bearer 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")
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求
	client := http.Client{}
	resp1, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp1.Body.Close()

	// 解析响应
	var respData map[string]interface{}
	err = json.NewDecoder(resp1.Body).Decode(&respData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印响应结果
	fmt.Println("Response:", respData)
	return nil, nil
}
