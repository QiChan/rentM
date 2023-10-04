package listen

import (
	"context"
	"rentM/app/usercenter/cmd/mq/internal/config"
	cronmq "rentM/app/usercenter/cmd/mq/internal/mqs/aq/cronMq"
	defermq "rentM/app/usercenter/cmd/mq/internal/mqs/aq/deferMq"
	"rentM/app/usercenter/cmd/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

func AsynqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	//HandlerProcessor := &defermq.HandlerProcessor{Ctx: ctx, SvcCtx: svcContext}
	return []service.Service{
		//监听延时队列
		/*
			defermq.NewAsynqDeferMq(ctx, svcContext, map[string]func(context.Context, *asynq.Task) error{
				defermsg.TypeFirstAsynqDeferTask:  HandlerProcessor.FirstDeferTaskHandler,
				defermsg.TypeSecondAsynqDeferTask: HandlerProcessor.SecondDeferTaskHandler,
			}),
		*/
		defermq.NewAsynqDeferMq(ctx, svcContext),

		//监听定时任务
		cronmq.NewAsynqCronMq(ctx, svcContext),
	}
}
