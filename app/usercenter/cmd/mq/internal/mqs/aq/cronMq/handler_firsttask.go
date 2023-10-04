package cronmq

import (
	"context"
	"encoding/json"
	"fmt"
	cronmsg "rentM/common/asynq_mq/cron_msg"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func (l *AsynqCronMq) FirstCronTaskHandler(ctx context.Context, t *asynq.Task) error {
	lock := redis.NewRedisLock(l.svcCtx.RedisCli, "firstCronTaskLock")
	lock.SetExpire(5)
	accquire, err := lock.Acquire()
	switch {
	case err != nil:
		fmt.Printf("first cron task get lock err: %v", err)
		lock.Release()
		return err
	case !accquire:
		fmt.Printf("first cron task get lock fail")
		lock.Release()
	case accquire:
		var p cronmsg.FirstCronTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		fmt.Printf("first cron task dealing succ! payload is %v", p)
		lock.Release()
	}
	return nil
}
