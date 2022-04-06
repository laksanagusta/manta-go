package taskUser

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(taskUser TaskUser) (TaskUser, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(taskUser TaskUser) (TaskUser, error) {
	err := r.db.Create(&taskUser).Error
	if err != nil {
		return taskUser, err
	}

	return taskUser, nil
}
