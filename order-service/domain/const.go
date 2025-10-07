package domain

type Status string

const (
	StatusPending  Status = "PENDING"
	StatusComplete Status = "COMPLETE"
	StatusFailed   Status = "FAILED"
)
