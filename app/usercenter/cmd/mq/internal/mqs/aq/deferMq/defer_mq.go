package defermq

import (
	"context"
	"fmt"
	"log"
	"rentM/app/usercenter/cmd/mq/internal/svc"
	defermsg "rentM/common/asynq_mq/defer_msg"

	"github.com/hibiken/asynq"
)

/*
type HandlerProcessor struct {
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}
*/

type AsynqDeferMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//typeHandlers map[string](func(context.Context, *asynq.Task) error)
}

// func NewAsynqDeferMq(ctx context.Context, svcCtx *svc.ServiceContext, typeHandlers map[string](func(context.Context, *asynq.Task) error)) *AsynqDeferMq {
func NewAsynqDeferMq(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqDeferMq {
	return &AsynqDeferMq{
		ctx:    ctx,
		svcCtx: svcCtx,
		//typeHandlers: typeHandlers,
	}
}

func (l *AsynqDeferMq) Start() {
	fmt.Println("AsynqDeferMq start ")

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
	mux.HandleFunc(defermsg.TypeFirstAsynqDeferTask, l.FirstDeferTaskHandler)
	mux.HandleFunc(defermsg.TypeSecondAsynqDeferTask, l.SecondDeferTaskHandler)
	/*
		for patt, fn := range l.typeHandlers {
			mux.HandleFunc(patt, fn)
		}
	*/

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func (l *AsynqDeferMq) Stop() {
	fmt.Println("AsynqDeferMq stop")
}
