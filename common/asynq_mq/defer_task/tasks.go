package defertask

import (
	"encoding/json"
	defermsg "rentM/common/asynq_mq/defer_msg"

	"github.com/hibiken/asynq"
)

func NewFirstAsynqDeferTask(uid int64, nickname string) (*asynq.Task, error) {
	payload, err := json.Marshal(defermsg.FirstDeferTaskPayload{UserID: uid, Nickname: nickname})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(defermsg.TypeFirstAsynqDeferTask, payload), nil
}

func NewSecondAsynqDeferTask(token string) (*asynq.Task, error) {
	payload, err := json.Marshal(defermsg.SecondDeferTaskPayload{Token: token})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(defermsg.TypeSecondAsynqDeferTask, payload), nil
}
