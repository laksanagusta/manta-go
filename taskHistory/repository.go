package taskHistory

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(taskHistory TaskHistory) (TaskHistory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

type Query map[string]interface{}

func (r *repository) Save(taskHistory TaskHistory) (TaskHistory, error) {
	err := r.db.Create(&taskHistory).Error
	if err != nil {
		return taskHistory, err
	}

	return taskHistory, nil
}
