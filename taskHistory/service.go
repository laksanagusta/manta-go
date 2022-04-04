package taskHistory

type Service interface {
	CreateTaskHistory(input CreateTaskHistoryInput) (TaskHistory, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTaskHistory(input CreateTaskHistoryInput) (TaskHistory, error) {
	taskHistory := TaskHistory{}

	newTaskHistory, err := s.repository.Save(taskHistory)
	if err != nil {
		return newTaskHistory, err
	}

	return newTaskHistory, nil
}
