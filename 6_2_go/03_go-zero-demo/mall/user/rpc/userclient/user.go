// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"go-zero-demo/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest         = user.IdRequest
	UserMoneyResponse = user.UserMoneyResponse
	UserResponse      = user.UserResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
		GetMoney(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserMoneyResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) GetMoney(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserMoneyResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetMoney(ctx, in, opts...)
}
