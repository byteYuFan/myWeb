// Code generated by Kitex v0.5.2. DO NOT EDIT.

package chatservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	chatgpt "myWeb/kitex_gen/chatgpt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Chat(ctx context.Context, Req *chatgpt.ChatRequest, callOptions ...callopt.Option) (r *chatgpt.ChatResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kChatServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kChatServiceClient struct {
	*kClient
}

func (p *kChatServiceClient) Chat(ctx context.Context, Req *chatgpt.ChatRequest, callOptions ...callopt.Option) (r *chatgpt.ChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Chat(ctx, Req)
}