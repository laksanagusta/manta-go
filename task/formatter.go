package task

import (
	"time"
)

type TaskFormatter struct {
	ID                int       `json:"id"`
	TaskTitle         string    `json:"taskTitle"`
	TaskSubTitle      string    `json:"taskSubtitle"`
	TaskType          string    `json:"taskType"`
	CustomerName      string    `json:"customerName"`
	TaskStartTime     time.Time `json:"taskStartTime"`
	TaskCompletedTime time.Time `json:"taskCompletedTime"`
}

func FormatTask(task Task) TaskFormatter {
	taskFormatter := TaskFormatter{}
	taskFormatter.ID = task.ID
	taskFormatter.TaskTitle = task.TaskTitle
	taskFormatter.TaskSubTitle = task.TaskSubTitle
	taskFormatter.TaskType = task.TaskType
	taskFormatter.CustomerName = task.CustomerName
	taskFormatter.TaskStartTime = task.TaskStartTime
	taskFormatter.TaskCompletedTime = task.TaskCompletedTime

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
