package task

import (
	"errors"
	"novocaine-dev/helper"
	"novocaine-dev/taskHistory"
	"novocaine-dev/taskUser"
	"time"
)

type Service interface {
	CreateTask(input CreateTaskInput) (Task, error)
	UpdateTask(input UpdateTaskInput) (Task, error)
	FindTaskById(id int) (Task, error)
	CustomFilter(query map[string][]string) ([]Task, error)
	Delete(id FindTaskById) (string, error)
	ProcessTask(input ProcessTaskInput, username string) ([]Task, error)
	AssignTask(input AssignTaskInput) (taskUser.TaskUser, error)
}

type service struct {
	repository      Repository
	taskHistoryRepo taskHistory.Repository
	taskUserRepo    taskUser.Repository
}

func NewService(repository Repository, taskHistoryRepo taskHistory.Repository, taskUserRepo taskUser.Repository) *service {
	return &service{repository, taskHistoryRepo, taskUserRepo}
}

func (s *service) CreateTask(input CreateTaskInput) (Task, error) {
	task := Task{}
	task.TaskTitle = input.TaskTitle
	task.TaskSubTitle = input.TaskSubTitle
	task.TaskType = input.TaskType
	task.CustomerName = input.CustomerName
	task.TaskRefId = input.TaskRefId
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
	task.TaskRefId = helper.CheckIfVariableSetsTypeString(task.TaskRefId, input.TaskRefId)
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

func (s *service) ProcessTask(input ProcessTaskInput, username string) ([]Task, error) {
	var content string = "Task processed " + input.Process + " by " + username
	var taskHistories taskHistory.TaskHistory

	taskHistories.Content = content
	taskHistories.User_id = input.UserIdProcessed

	tasks, err := s.repository.GetTaskByMultipleRefId(input.Tasks)
	if err != nil {
		return tasks, err
	}

	for _, v := range tasks {
		taskHistories.Task_id = v.ID
		_, err := s.taskHistoryRepo.Save(taskHistories)
		if err != nil {
			return tasks, err
		}
	}

	return tasks, nil
}

func (s *service) AssignTask(input AssignTaskInput) (taskUser.TaskUser, error) {
	var taskUser taskUser.TaskUser

	taskUser.Task_id = input.TaskID
	taskUser.User_id = input.UserID

	newTaskUser, err := s.taskUserRepo.Save(taskUser)
	if err != nil {
		return newTaskUser, err
	}

	return newTaskUser, nil
}
