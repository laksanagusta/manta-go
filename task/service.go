package task

import (
	"errors"
	"novocaine-dev/helper"
	"time"
)

type Service interface {
	CreateTask(input CreateTaskInput) (Task, error)
	UpdateTask(input UpdateTaskInput) (Task, error)
	FindTaskById(id int) (Task, error)
	CustomFilter(query map[string][]string) ([]Task, error)
	Delete(id FindTaskById) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTask(input CreateTaskInput) (Task, error) {
	task := Task{}
	task.TaskTitle = input.TaskTitle
	task.TaskSubTitle = input.TaskSubTitle
	task.TaskType = input.TaskType
	task.CustomerName = input.CustomerName
	task.TaskStartTime = time.Time{}
	task.TaskCompletedTime = time.Time{}

	newTask, err := s.repository.Save(task)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

func (s *service) UpdateTask(input UpdateTaskInput) (Task, error) {

	task, err := s.repository.FindTaskById(input.ID)
	if err != nil {
		return task, err
	}

	task.TaskTitle = helper.CheckIfVariableSetsTypeString(task.TaskTitle, input.TaskTitle)
	task.TaskSubTitle = input.TaskSubTitle
	task.TaskType = helper.CheckIfVariableSetsTypeString(task.TaskType, input.TaskType)
	task.CustomerName = helper.CheckIfVariableSetsTypeString(task.CustomerName, input.CustomerName)
	task.TaskStartTime = time.Time{}
	task.TaskCompletedTime = time.Time{}

	updateTask, err := s.repository.Update(task)
	if err != nil {
		return updateTask, err
	}

	return updateTask, nil
}

func (s *service) FindTaskById(id int) (Task, error) {
	task, err := s.repository.FindTaskById(id)
	if err != nil {
		return task, err
	}

	if task.ID == 0 {
		return task, errors.New("Task not found")
	}

	return task, nil
}

func (s *service) CustomFilter(query map[string][]string) ([]Task, error) {
	task, err := s.repository.CustomFilter(query)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *service) Delete(id FindTaskById) (string, error) {
	task, err := s.repository.Delete(id.ID)
	if err != nil {
		return task, err
	}

	return task, nil
}
