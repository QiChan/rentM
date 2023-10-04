package user

import (
	"context"

	"rentM/app/usercenter/cmd/api/internal/svc"
	"rentM/app/usercenter/cmd/api/internal/types"
	"rentM/app/usercenter/cmd/rpc/pb"
	"rentM/common/ctxdata"
	structurecheck "rentM/common/structure_check"
	"rentM/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoModLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoModLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoModLogic {
	return &UserInfoModLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoModLogic) UserInfoMod(req *types.UserInfoModReq) (resp *types.UserInfoModResp, err error) {
	// todo: add your logic here and delete this line
	if empty := structurecheck.CheckField(*req); empty {
		return nil, errors.Wrapf(xerr.NewErrMsg("empty req"), "need at least one member not empty")
	}

	userId := ctxdata.GetUidFromCtx(l.ctx)
	modResp, err := l.svcCtx.UserCenterRpc.UserInfoMod(l.ctx, &pb.UserInfoModReq{
		Uid:      userId,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Info:     req.Info,
		Password: req.Password,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var rsp types.UserInfoModResp
	_ = copier.Copy(&rsp, modResp)

	return &rsp, nil
}
