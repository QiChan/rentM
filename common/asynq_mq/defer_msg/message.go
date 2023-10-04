package defermsg

const (
	TypeFirstAsynqDeferTask  = "defer:one"
	TypeSecondAsynqDeferTask = "defer:two"
)

type FirstDeferTaskPayload struct {
	UserID   int64
	Nickname string
}

type SecondDeferTaskPayload struct {
	Token string
}
