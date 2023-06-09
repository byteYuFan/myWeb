// Code generated by Kitex v0.5.2. DO NOT EDIT.

package chatservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	chatgpt "myWeb/kitex_gen/chatgpt"
)

func serviceInfo() *kitex.ServiceInfo {
	return chatServiceServiceInfo
}

var chatServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ChatService"
	handlerType := (*chatgpt.ChatService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Chat": kitex.NewMethodInfo(chatHandler, newChatArgs, newChatResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "chatgpt",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func chatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chatgpt.ChatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chatgpt.ChatService).Chat(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChatArgs:
		success, err := handler.(chatgpt.ChatService).Chat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChatResult)
		realResult.Success = success
	}
	return nil
}
func newChatArgs() interface{} {
	return &ChatArgs{}
}

func newChatResult() interface{} {
	return &ChatResult{}
}

type ChatArgs struct {
	Req *chatgpt.ChatRequest
}

func (p *ChatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chatgpt.ChatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChatArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChatArgs) Unmarshal(in []byte) error {
	msg := new(chatgpt.ChatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChatArgs_Req_DEFAULT *chatgpt.ChatRequest

func (p *ChatArgs) GetReq() *chatgpt.ChatRequest {
	if !p.IsSetReq() {
		return ChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChatResult struct {
	Success *chatgpt.ChatResponse
}

var ChatResult_Success_DEFAULT *chatgpt.ChatResponse

func (p *ChatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chatgpt.ChatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChatResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChatResult) Unmarshal(in []byte) error {
	msg := new(chatgpt.ChatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChatResult) GetSuccess() *chatgpt.ChatResponse {
	if !p.IsSetSuccess() {
		return ChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*chatgpt.ChatResponse)
}

func (p *ChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChatResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Chat(ctx context.Context, Req *chatgpt.ChatRequest) (r *chatgpt.ChatResponse, err error) {
	var _args ChatArgs
	_args.Req = Req
	var _result ChatResult
	if err = p.c.Call(ctx, "Chat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
