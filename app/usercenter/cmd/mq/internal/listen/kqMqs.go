package listen

import (
	"context"
	"rentM/app/usercenter/cmd/mq/internal/config"
	kqmq "rentM/app/usercenter/cmd/mq/internal/mqs/kq"
	"rentM/app/usercenter/cmd/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		kq.MustNewQueue(c.KafkaFirstTestConf, kqmq.NewFirstKqTest(ctx, svcContext)),
	}

}
