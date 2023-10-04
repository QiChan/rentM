package logic

import (
	"context"

	"rentM/app/usercenter/cmd/rpc/internal/svc"
	"rentM/app/usercenter/cmd/rpc/pb"
	"rentM/app/usercenter/model"
	"rentM/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoModLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoModLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoModLogic {
	return &UserInfoModLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoModLogic) UserInfoMod(in *pb.UserInfoModReq) (*pb.UserInfoModResp, error) {
	// todo: add your logic here and delete this line
	if userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Uid); err != nil {
		logx.WithContext(l.ctx).Error("call UserInfoMod() at 1 fail: ", err)
		return nil, err
	} else if userInfo != nil {
		if in.Password == "" {
			in.Password = userInfo.Password
		}
		if in.Nickname == "" {
			in.Nickname = userInfo.Nickname
		}
		if in.Avatar == "" {
			in.Avatar = userInfo.Avatar
		}
		if in.Info == "" {
			in.Info = userInfo.Info
		}
		if _, err := l.svcCtx.UserModel.CustomUpdate(l.ctx, nil, &model.User{
			Id: in.Uid,
		}, map[string]interface{}{
			"nickname": in.Nickname,
			"avatar":   in.Avatar,
			"info":     in.Info,
			"password": tool.Md5ByString(in.Password),
		}); err != nil {
			logx.WithContext(l.ctx).Error("call UserInfoMod() at 2 fail: ", err)
			return nil, err
		}
	}

	return &pb.UserInfoModResp{Ok: true}, nil
}
