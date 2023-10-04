package listen

import (
	"context"
	"fmt"
	"rentM/app/usercenter/cmd/mq/internal/config"
	"rentM/app/usercenter/cmd/mq/internal/svc"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/service"
)

func RcronSvc(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	//HandlerProcessor := &defermq.HandlerProcessor{Ctx: ctx, SvcCtx: svcContext}
	return []service.Service{
		NewRCron(ctx, svcContext),
	}
}

type Rcron struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRCron(ctx context.Context, svcCtx *svc.ServiceContext) *Rcron {
	return &Rcron{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Rcron) Start() {
	fmt.Println("Rcron run")
	supportSecParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	cr := cron.New(cron.WithParser(supportSecParser), cron.WithChain())
	cr.AddFunc("30 */1 * * * *", func() {
		fmt.Println("Rcron alarming u~")
	})

	cr.Run()
}

func (l *Rcron) Stop() {
	fmt.Println("Rcron stop")
}
