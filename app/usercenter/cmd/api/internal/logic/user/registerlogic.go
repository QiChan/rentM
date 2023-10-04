package user

import (
	"context"
	"fmt"
	"time"

	"rentM/app/usercenter/cmd/api/internal/svc"
	"rentM/app/usercenter/cmd/api/internal/types"
	"rentM/app/usercenter/cmd/rpc/pb"
	"rentM/app/usercenter/model"
	defertask "rentM/common/asynq_mq/defer_task"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserCenterRpc.Register(l.ctx, &pb.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
		MsgCode:  req.MsgCode,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var rsp types.RegisterResp
	_ = copier.Copy(&rsp, registerResp)

	//延时队列测试2
	task, errx := defertask.NewSecondAsynqDeferTask(registerResp.AccessToken)
	if errx != nil {
		fmt.Printf("second defer task create err: %v", errx)
	}
	if info, err := l.svcCtx.AqDeferCli.Enqueue(task, asynq.TaskID("secondDeferTask"), asynq.ProcessIn(30*time.Second)); err != nil {
		fmt.Printf("enqueue second defer task err: %v", err)
	} else {
		fmt.Printf("enqueue second defer task succ, id = %s, queue = %s", info.ID, info.Queue)
	}

	return &rsp, nil
}
