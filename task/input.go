package task

import (
	"time"
)

type FindTaskById struct {
	ID int `uri:"id" binding:required`
}

type CreateTaskInput struct {
	TaskTitle         string    `json:"taskTitle" binding:"required"`
	TaskSubTitle      string    `json:"taskSubTitle"`
	TaskType          string    `json:"taskType" binding:"required"`
	CustomerName      string    `json:"customerName"`
	TaskStartTime     time.Time `json:"taskStartTime"`
	TaskCompletedTime time.Time `json:"taskCompletedTime"`
}

type UpdateTaskInput struct {
	ID                int       `json:"id" binding:"required"`
	TaskTitle         string    `json:"taskTitle"`
	TaskSubTitle      string    `json:"taskSubTitle"`
	TaskType          string    `json:"taskType"`
	CustomerName      string    `json:"customerName"`
	TaskStartTime     time.Time `json:"taskStartTime"`
	TaskCompletedTime time.Time `json:"taskCompletedTime"`
}
