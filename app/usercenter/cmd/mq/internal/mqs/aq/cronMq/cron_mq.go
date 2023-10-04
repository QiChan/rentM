package cronmq

import (
	"context"
	"fmt"
	"log"
	"rentM/app/usercenter/cmd/mq/internal/svc"
	cronmsg "rentM/common/asynq_mq/cron_msg"

	"github.com/hibiken/asynq"
)

type AsynqCronMq struct {
	ctx    context.Context
	svcCtx svc.ServiceContext
}

func NewAsynqCronMq(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqCronMq {
	return &AsynqCronMq{
		ctx:    ctx,
		svcCtx: *svcCtx,
	}
}

func (l *AsynqCronMq) Start() {
	fmt.Println("AsynqCronMq start ")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()

	//注册主题和处理函数
	mux.HandleFunc(cronmsg.TypeFirstAsynqCronTask, l.FirstCronTaskHandler)
	/*
		for patt, fn := range l.typeHandlers {
			mux.HandleFunc(patt, fn)
		}
	*/

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func (l *AsynqCronMq) Stop() {
	fmt.Println("AsynqCronMq stop")
}
