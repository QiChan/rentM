package kq

import (
	"context"
	"encoding/json"
	"fmt"
	"rentM/app/usercenter/cmd/mq/internal/svc"
	"rentM/common/kqueue"

	"github.com/zeromicro/go-zero/core/logx"
)

type FirstKqTest struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFirstKqTest(ctx context.Context, svcCtx *svc.ServiceContext) *FirstKqTest {
	return &FirstKqTest{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FirstKqTest) Consume(key, val string) error {
	var message kqueue.KqueueFirstTestMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("FirstKqTest->Consume Unmarshal err : %v, key: %s ,val : %s", err, val)
		return err
	}

	fmt.Printf("first kq test succ! key is: %s, message is: %v", key, message)

	return nil
}
