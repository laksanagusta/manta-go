package task

import (
	"novocaine-dev/taskHistory"
	"novocaine-dev/user"
	"time"
)

type TaskFormatter struct {
	ID                int                       `json:"id"`
	TaskTitle         string                    `json:"taskTitle"`
	TaskSubTitle      string                    `json:"taskSubtitle"`
	TaskType          string                    `json:"taskType"`
	CustomerName      string                    `json:"customerName"`
	TaskHistory       []taskHistory.TaskHistory `json:"taskHistory"`
	Users             []user.User               `json:"users"`
	TaskRefId         string                    `json:"taskRefId"`
	TaskStartTime     time.Time                 `json:"taskStartTime"`
	TaskCompletedTime time.Time                 `json:"taskCompletedTime"`
	CreatedAt         time.Time                 `json:"createdAt"`
	UpdatedAt         time.Time                 `json:"updatedAt"`
}

func FormatTask(task Task) TaskFormatter {
	taskFormatter := TaskFormatter{}
	taskFormatter.ID = task.ID
	taskFormatter.TaskTitle = task.TaskTitle
	taskFormatter.TaskSubTitle = task.TaskSubTitle
	taskFormatter.TaskType = task.TaskType
	taskFormatter.CustomerName = task.CustomerName
	taskFormatter.TaskHistory = task.TaskHistories
	taskFormatter.Users = task.Users
	taskFormatter.TaskRefId = task.TaskRefId
	taskFormatter.TaskStartTime = task.TaskStartTime
	taskFormatter.TaskCompletedTime = task.TaskCompletedTime
	taskFormatter.CreatedAt = task.CreatedAt
	taskFormatter.UpdatedAt = task.UpdatedAt

	return taskFormatter
}

func FormatTasks(tasks []Task) []TaskFormatter {
	tasksFormatter := []TaskFormatter{}

	for _, task := range tasks {
		taskFormatter := FormatTask(task)
		tasksFormatter = append(tasksFormatter, taskFormatter)
	}

	return tasksFormatter
}
