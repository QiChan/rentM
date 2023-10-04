// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"

	"rentM/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq         = pb.GenerateTokenReq
	GenerateTokenResp        = pb.GenerateTokenResp
	GetUserAuthByAuthKeyReq  = pb.GetUserAuthByAuthKeyReq
	GetUserAuthByAuthKeyResp = pb.GetUserAuthByAuthKeyResp
	GetUserAuthByUserIdReq   = pb.GetUserAuthByUserIdReq
	GetUserAuthyUserIdResp   = pb.GetUserAuthyUserIdResp
	GetUserInfoReq           = pb.GetUserInfoReq
	GetUserInfoResp          = pb.GetUserInfoResp
	LoginReq                 = pb.LoginReq
	LoginResp                = pb.LoginResp
	RegisterReq              = pb.RegisterReq
	RegisterResp             = pb.RegisterResp
	User                     = pb.User
	UserAuth                 = pb.UserAuth
	UserInfoModReq           = pb.UserInfoModReq
	UserInfoModResp          = pb.UserInfoModResp

	UserCenter interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error)
		GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		UserInfoMod(ctx context.Context, in *UserInfoModReq, opts ...grpc.CallOption) (*UserInfoModResp, error)
	}

	defaultUserCenter struct {
		cli zrpc.Client
	}
)

func NewUserCenter(cli zrpc.Client) UserCenter {
	return &defaultUserCenter{
		cli: cli,
	}
}

func (m *defaultUserCenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserCenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserCenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserCenter) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.GetUserAuthByAuthKey(ctx, in, opts...)
}

func (m *defaultUserCenter) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.GetUserAuthByUserId(ctx, in, opts...)
}

func (m *defaultUserCenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUserCenter) UserInfoMod(ctx context.Context, in *UserInfoModReq, opts ...grpc.CallOption) (*UserInfoModResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.UserInfoMod(ctx, in, opts...)
}
