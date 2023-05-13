// Code generated by Kitex v0.5.2. DO NOT EDIT.

package registerservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	usersrv "myWeb/kitex_gen/usersrv"
)

func serviceInfo() *kitex.ServiceInfo {
	return registerServiceServiceInfo
}

var registerServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RegisterService"
	handlerType := (*usersrv.RegisterService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register": kitex.NewMethodInfo(registerHandler, newRegisterServiceRegisterArgs, newRegisterServiceRegisterResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "usersrv",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersrv.RegisterServiceRegisterArgs)
	realResult := result.(*usersrv.RegisterServiceRegisterResult)
	success, err := handler.(usersrv.RegisterService).Register(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRegisterServiceRegisterArgs() interface{} {
	return usersrv.NewRegisterServiceRegisterArgs()
}

func newRegisterServiceRegisterResult() interface{} {
	return usersrv.NewRegisterServiceRegisterResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, request *usersrv.RegisterRequest) (r *usersrv.RegisterResponse, err error) {
	var _args usersrv.RegisterServiceRegisterArgs
	_args.Request = request
	var _result usersrv.RegisterServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
