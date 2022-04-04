package task

import (
	"time"
)

type Task struct {
	ID                int
	TaskTitle         string
	TaskSubTitle      string
	TaskType          string
	CustomerName      string
	TaskCreatedBy     int
	TaskAssignedTo    int
	TaskStartTime     time.Time
	TaskCompletedTime time.Time
}
