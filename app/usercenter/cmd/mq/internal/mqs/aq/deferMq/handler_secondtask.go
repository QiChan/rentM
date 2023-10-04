package defermq

import (
	"context"
	"encoding/json"
	"fmt"

	defermsg "rentM/common/asynq_mq/defer_msg"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func (l *AsynqDeferMq) SecondDeferTaskHandler(ctx context.Context, t *asynq.Task) error {
	lock := redis.NewRedisLock(l.svcCtx.RedisCli, "secondDeferTaskLock")
	lock.SetExpire(5)
	accquire, err := lock.Acquire()
	switch {
	case err != nil:
		fmt.Printf("second defer task get lock err: %v", err)
		lock.Release()
		return err
	case !accquire:
		fmt.Printf("second defer task get lock fail")
		lock.Release()
	case accquire:
		var p defermsg.SecondDeferTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		fmt.Printf("second defer task dealing succ! payload is %v", p)
		lock.Release()
	}
	return nil
}
