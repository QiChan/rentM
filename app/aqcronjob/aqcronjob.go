package main

import (
	"encoding/json"
	"fmt"
	"log"

	cronmsg "rentM/common/asynq_mq/cron_msg"

	"github.com/hibiken/asynq"
)

const redisAddr = "192.168.1.21:6379"
const redisPwd = "alp821821"

func main() {
	// 周期性任务
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Password: redisPwd,
		}, nil)

	payload, err := json.Marshal(cronmsg.FirstCronTaskPayload{Email: "497267655@qq.com", Content: "该发爱的邮件给包包喇~"})
	if err != nil {
		log.Fatal(err)
	}

	task := asynq.NewTask(cronmsg.TypeFirstAsynqCronTask, payload, asynq.TaskID("firstCronTask"))
	// 每隔1分钟同步一次
	entryID, err := scheduler.Register("*/1 * * * *", task)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("registered an entry: %q\n", entryID)

	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
