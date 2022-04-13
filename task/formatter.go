package task

import (
	"time"
)

type TaskFormatter struct {
	ID                int                         `json:"id"`
	TaskTitle         string                      `json:"taskTitle"`
	TaskSubTitle      string                      `json:"taskSubtitle"`
	TaskType          string                      `json:"taskType"`
	CustomerName      string                      `json:"customerName"`
	TaskHistory       []TaskHistoriesFormatterOut `json:"taskHistory"`
	Users             []TaskUsersFormatterOut     `json:"users"`
	TaskRefId         string                      `json:"taskRefId"`
	TaskStartTime     time.Time                   `json:"taskStartTime"`
	TaskCompletedTime time.Time                   `json:"taskCompletedTime"`
	CreatedAt         time.Time                   `json:"createdAt"`
	UpdatedAt         time.Time                   `json:"updatedAt"`
}

type TaskUsersFormatterOut struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	AssignedAs string    `json:"assignedAs"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type TaskHistoriesFormatterOut struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	LatestStatus string `json:"latestStatus"`
}

func FormatTask(task Task) TaskFormatter {
	taskFormatter := TaskFormatter{}
	taskFormatter.ID = task.ID
	taskFormatter.TaskTitle = task.TaskTitle
	taskFormatter.TaskSubTitle = task.TaskSubTitle
	taskFormatter.TaskType = task.TaskType
	taskFormatter.CustomerName = task.CustomerName

	users := []TaskUsersFormatterOut{}
	for _, v := range task.Users {
		taskUsersFormatterOut := TaskUsersFormatterOut{}
		taskUsersFormatterOut.ID = v.ID
		taskUsersFormatterOut.Name = v.Name
		taskUsersFormatterOut.Email = v.Email
		taskUsersFormatterOut.Role = v.Role
		taskUsersFormatterOut.Username = v.Username
		taskUsersFormatterOut.CreatedAt = v.CreatedAt
		users = append(users, taskUsersFormatterOut)
	}

	taskFormatter.Users = users

	histories := []TaskHistoriesFormatterOut{}
	for _, v := range task.TaskHistories {
		taskHistoriesFormatterOut := TaskHistoriesFormatterOut{}
		taskHistoriesFormatterOut.ID = v.ID
		taskHistoriesFormatterOut.Content = v.Content
		taskHistoriesFormatterOut.LatestStatus = v.LatestStatus
		histories = append(histories, taskHistoriesFormatterOut)
	}

	taskFormatter.TaskHistory = histories

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
