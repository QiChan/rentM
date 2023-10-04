package listen

import (
	"context"
	"rentM/app/usercenter/cmd/mq/internal/config"
	"rentM/app/usercenter/cmd/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq ：pub sub
	services = append(services, KqMqs(c, ctx, svcContext)...)
	//aq
	services = append(services, AsynqMqs(c, ctx, svcContext)...)
	//robfig cron
	services = append(services, RcronSvc(c, ctx, svcContext)...)

	return services
}
