// Code generated by Kitex v0.5.2. DO NOT EDIT.
package userservice

import (
	demouser "github.com/yiwen101/CardWizards/kitex_gen/demouser"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler demouser.UserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
