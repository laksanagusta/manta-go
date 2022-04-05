package task

import (
	"novocaine-dev/taskHistory"
	"novocaine-dev/user"
	"time"
)

type Task struct {
	ID                int
	TaskTitle         string
	TaskSubTitle      string
	TaskType          string
	CustomerName      string
	TaskHistories     []taskHistory.TaskHistory
	Users             []user.User `gorm:"many2many:task_users"`
	TaskStartTime     time.Time
	TaskCompletedTime time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
