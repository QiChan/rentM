package cronmsg

const (
	TypeFirstAsynqCronTask = "cron:one"
)

type FirstCronTaskPayload struct {
	Email   string
	Content string
}
