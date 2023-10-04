package user

import (
	"context"
	"encoding/json"
	"strconv"

	"rentM/app/usercenter/cmd/api/internal/svc"
	"rentM/app/usercenter/cmd/api/internal/types"
	"rentM/app/usercenter/cmd/rpc/pb"
	"rentM/app/usercenter/model"
	"rentM/common/kqueue"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.UserCenterRpc.Login(l.ctx, &pb.LoginReq{
		AuthType: model.UserAuthTypeSystem,
		AuthKey:  req.Mobile,
		Password: req.Password,
		MsgCode:  req.MsgCode,
	})
	if err != nil {
		return nil, err
	}

	var rsp types.LoginResp
	_ = copier.Copy(&rsp, loginResp)

	//测试kafka，每次登录成功发msg给kafka处理
	phoneNumber, _ := strconv.ParseInt(req.Mobile, 10, 64)
	data, _ := json.Marshal(
		kqueue.KqueueFirstTestMessage{
			Mobile: phoneNumber,
			Token:  rsp.AccessToken,
		})
	if err := l.svcCtx.KqPusherClient.Push(string(data)); err != nil {
		logx.Errorf("KqPusherClient Push err: %v", err)
	}

	return &rsp, nil
}
