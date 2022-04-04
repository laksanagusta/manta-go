package taskHistory

import "time"

type TaskHistory struct {
	ID           int
	Content      string
	LatestStatus string
	User_id      int
	Task_id      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
