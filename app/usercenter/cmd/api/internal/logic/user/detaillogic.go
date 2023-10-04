package user

import (
	"context"
	"fmt"
	"time"

	"rentM/app/usercenter/cmd/api/internal/svc"
	"rentM/app/usercenter/cmd/api/internal/types"
	"rentM/app/usercenter/cmd/rpc/pb"
	defertask "rentM/common/asynq_mq/defer_task"
	"rentM/common/ctxdata"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)

	userInfoResp, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	var userInfo types.User
	_ = copier.Copy(&userInfo, userInfoResp.User)

	//延时队列测试1
	task, errx := defertask.NewFirstAsynqDeferTask(userId, userInfoResp.User.Nickname)
	if errx != nil {
		fmt.Printf("first defer task create err: %v", errx)
	}
	if info, err := l.svcCtx.AqDeferCli.Enqueue(task, asynq.TaskID("firstDeferTask"), asynq.ProcessIn(30*time.Second)); err != nil {
		fmt.Printf("enqueue first defer task err: %v", err)
	} else {
		fmt.Printf("enqueue first defer task succ, id = %s, queue = %s", info.ID, info.Queue)
	}

	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
