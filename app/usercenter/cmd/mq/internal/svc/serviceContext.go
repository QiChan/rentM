package svc

import (
	"rentM/app/usercenter/cmd/mq/internal/config"
	"rentM/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	RedisCli *redis.Redis

	UsercenterRpc usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//RedisCli: redis.MustNewRedis(redis.RedisConf{Host: "192.168.1.21:6379", Type: "node", Pass: "alp821821"}),
		RedisCli: redis.MustNewRedis(redis.RedisConf{Host: c.Redis.Host, Type: c.Redis.Type, Pass: c.Redis.Pass}),

		UsercenterRpc: usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
