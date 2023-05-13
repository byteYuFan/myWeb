// Code generated by Kitex v0.5.2. DO NOT EDIT.
package registerservice

import (
	server "github.com/cloudwego/kitex/server"
	usersrv "myWeb/kitex_gen/user/usersrv"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler usersrv.RegisterService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
