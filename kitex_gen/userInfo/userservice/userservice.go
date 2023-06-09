// Code generated by Kitex v0.5.2. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	userInfo "myWeb/kitex_gen/userInfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userInfo.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UpdateUser":     kitex.NewMethodInfo(updateUserHandler, newUpdateUserArgs, newUpdateUserResult, false),
		"ChangePassword": kitex.NewMethodInfo(changePasswordHandler, newChangePasswordArgs, newChangePasswordResult, false),
		"GetUser":        kitex.NewMethodInfo(getUserHandler, newGetUserArgs, newGetUserResult, false),
		"RestPassword":   kitex.NewMethodInfo(restPasswordHandler, newRestPasswordArgs, newRestPasswordResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "userInfo",
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

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userInfo.UpdateUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userInfo.UserService).UpdateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateUserArgs:
		success, err := handler.(userInfo.UserService).UpdateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateUserArgs() interface{} {
	return &UpdateUserArgs{}
}

func newUpdateUserResult() interface{} {
	return &UpdateUserResult{}
}

type UpdateUserArgs struct {
	Req *userInfo.UpdateUserRequest
}

func (p *UpdateUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userInfo.UpdateUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserArgs) Unmarshal(in []byte) error {
	msg := new(userInfo.UpdateUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserArgs_Req_DEFAULT *userInfo.UpdateUserRequest

func (p *UpdateUserArgs) GetReq() *userInfo.UpdateUserRequest {
	if !p.IsSetReq() {
		return UpdateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserResult struct {
	Success *userInfo.UpdateUserResponse
}

var UpdateUserResult_Success_DEFAULT *userInfo.UpdateUserResponse

func (p *UpdateUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userInfo.UpdateUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserResult) Unmarshal(in []byte) error {
	msg := new(userInfo.UpdateUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserResult) GetSuccess() *userInfo.UpdateUserResponse {
	if !p.IsSetSuccess() {
		return UpdateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userInfo.UpdateUserResponse)
}

func (p *UpdateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserResult) GetResult() interface{} {
	return p.Success
}

func changePasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userInfo.ChangePasswordRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userInfo.UserService).ChangePassword(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChangePasswordArgs:
		success, err := handler.(userInfo.UserService).ChangePassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChangePasswordResult)
		realResult.Success = success
	}
	return nil
}
func newChangePasswordArgs() interface{} {
	return &ChangePasswordArgs{}
}

func newChangePasswordResult() interface{} {
	return &ChangePasswordResult{}
}

type ChangePasswordArgs struct {
	Req *userInfo.ChangePasswordRequest
}

func (p *ChangePasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userInfo.ChangePasswordRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChangePasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChangePasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChangePasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChangePasswordArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChangePasswordArgs) Unmarshal(in []byte) error {
	msg := new(userInfo.ChangePasswordRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChangePasswordArgs_Req_DEFAULT *userInfo.ChangePasswordRequest

func (p *ChangePasswordArgs) GetReq() *userInfo.ChangePasswordRequest {
	if !p.IsSetReq() {
		return ChangePasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChangePasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChangePasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChangePasswordResult struct {
	Success *userInfo.ChangePasswordResponse
}

var ChangePasswordResult_Success_DEFAULT *userInfo.ChangePasswordResponse

func (p *ChangePasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userInfo.ChangePasswordResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChangePasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChangePasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChangePasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChangePasswordResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChangePasswordResult) Unmarshal(in []byte) error {
	msg := new(userInfo.ChangePasswordResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangePasswordResult) GetSuccess() *userInfo.ChangePasswordResponse {
	if !p.IsSetSuccess() {
		return ChangePasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChangePasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*userInfo.ChangePasswordResponse)
}

func (p *ChangePasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChangePasswordResult) GetResult() interface{} {
	return p.Success
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userInfo.GetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userInfo.UserService).GetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserArgs:
		success, err := handler.(userInfo.UserService).GetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserArgs() interface{} {
	return &GetUserArgs{}
}

func newGetUserResult() interface{} {
	return &GetUserResult{}
}

type GetUserArgs struct {
	Req *userInfo.GetUserRequest
}

func (p *GetUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userInfo.GetUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserArgs) Unmarshal(in []byte) error {
	msg := new(userInfo.GetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserArgs_Req_DEFAULT *userInfo.GetUserRequest

func (p *GetUserArgs) GetReq() *userInfo.GetUserRequest {
	if !p.IsSetReq() {
		return GetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserResult struct {
	Success *userInfo.GetUserResponse
}

var GetUserResult_Success_DEFAULT *userInfo.GetUserResponse

func (p *GetUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userInfo.GetUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserResult) Unmarshal(in []byte) error {
	msg := new(userInfo.GetUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserResult) GetSuccess() *userInfo.GetUserResponse {
	if !p.IsSetSuccess() {
		return GetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userInfo.GetUserResponse)
}

func (p *GetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserResult) GetResult() interface{} {
	return p.Success
}

func restPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userInfo.ResetPasswordRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userInfo.UserService).RestPassword(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RestPasswordArgs:
		success, err := handler.(userInfo.UserService).RestPassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RestPasswordResult)
		realResult.Success = success
	}
	return nil
}
func newRestPasswordArgs() interface{} {
	return &RestPasswordArgs{}
}

func newRestPasswordResult() interface{} {
	return &RestPasswordResult{}
}

type RestPasswordArgs struct {
	Req *userInfo.ResetPasswordRequest
}

func (p *RestPasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userInfo.ResetPasswordRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RestPasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RestPasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RestPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RestPasswordArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RestPasswordArgs) Unmarshal(in []byte) error {
	msg := new(userInfo.ResetPasswordRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RestPasswordArgs_Req_DEFAULT *userInfo.ResetPasswordRequest

func (p *RestPasswordArgs) GetReq() *userInfo.ResetPasswordRequest {
	if !p.IsSetReq() {
		return RestPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RestPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RestPasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RestPasswordResult struct {
	Success *userInfo.ResetPasswordResponse
}

var RestPasswordResult_Success_DEFAULT *userInfo.ResetPasswordResponse

func (p *RestPasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userInfo.ResetPasswordResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RestPasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RestPasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RestPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RestPasswordResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RestPasswordResult) Unmarshal(in []byte) error {
	msg := new(userInfo.ResetPasswordResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RestPasswordResult) GetSuccess() *userInfo.ResetPasswordResponse {
	if !p.IsSetSuccess() {
		return RestPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RestPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*userInfo.ResetPasswordResponse)
}

func (p *RestPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RestPasswordResult) GetResult() interface{} {
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

func (p *kClient) UpdateUser(ctx context.Context, Req *userInfo.UpdateUserRequest) (r *userInfo.UpdateUserResponse, err error) {
	var _args UpdateUserArgs
	_args.Req = Req
	var _result UpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangePassword(ctx context.Context, Req *userInfo.ChangePasswordRequest) (r *userInfo.ChangePasswordResponse, err error) {
	var _args ChangePasswordArgs
	_args.Req = Req
	var _result ChangePasswordResult
	if err = p.c.Call(ctx, "ChangePassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, Req *userInfo.GetUserRequest) (r *userInfo.GetUserResponse, err error) {
	var _args GetUserArgs
	_args.Req = Req
	var _result GetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RestPassword(ctx context.Context, Req *userInfo.ResetPasswordRequest) (r *userInfo.ResetPasswordResponse, err error) {
	var _args RestPasswordArgs
	_args.Req = Req
	var _result RestPasswordResult
	if err = p.c.Call(ctx, "RestPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
