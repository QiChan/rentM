package svc

import (
	"rentM/app/usercenter/cmd/api/internal/config"
	"rentM/app/usercenter/cmd/rpc/usercenter"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserCenterRpc  usercenter.UserCenter
	KqPusherClient *kq.Pusher
	RedisCli       *redis.Redis
	AqDeferCli     *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserCenterRpc:  usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		RedisCli:       redis.MustNewRedis(c.Redis),
		AqDeferCli:     asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass}),
	}
}
