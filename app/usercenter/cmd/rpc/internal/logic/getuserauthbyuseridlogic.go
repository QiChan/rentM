package logic

import (
	"context"

	"rentM/app/usercenter/cmd/rpc/internal/svc"
	"rentM/app/usercenter/cmd/rpc/pb"
	"rentM/app/usercenter/cmd/rpc/usercenter"
	"rentM/app/usercenter/model"
	"rentM/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *pb.GetUserAuthByUserIdReq) (*pb.GetUserAuthyUserIdResp, error) {
	// todo: add your logic here and delete this line
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeUserId(l.ctx, in.AuthType, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("获取用户授权信息失败"), "err : %v , in : %+v", err, in)
	}

	//解析返回数据.
	var respUserAuth usercenter.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)

	return &pb.GetUserAuthyUserIdResp{
		UserAuth: &respUserAuth,
	}, nil
}
