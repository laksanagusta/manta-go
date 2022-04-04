package task

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type Repository interface {
	FindTaskById(id int) (Task, error)
	Save(task Task) (Task, error)
	Update(task Task) (Task, error)
	CustomFilter(query map[string][]string) ([]Task, error)
	Delete(id int) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

type Query map[string]interface{}

func (r *repository) FindTaskById(id int) (Task, error) {
	var task Task
	err := r.db.Where("id = ?", id).Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Save(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Update(task Task) (Task, error) {
	err := r.db.Updates(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) CustomFilter(query map[string][]string) ([]Task, error) {
	var task []Task

	if len(query) == 0 {
		err := r.db.Find(&task).Limit(20).Error

		if err != nil {
			return task, err
		}
	}

	var i int = 1
	for k, v := range query {
		if v[0] != "" {
			if i == len(query) {
				if k == "limit" {
					limit, _ := strconv.Atoi(v[0])
					err := r.db.Limit(limit).Find(&task).Error
					if err != nil {
						return task, err
					}
					continue
				}

				err := r.db.Where(k+" = ?", v[0]).Limit(20).Find(&task).Error
				if err != nil {
					return task, err
				}

			} else {
				r.db.Where(k+" = ?", v[0])
			}
		}

		i++
	}

	return task, nil
}

func (r *repository) Delete(id int) (string, error) {
	deleteTaskDb := r.db.Delete(&Task{}, 28)
	if deleteTaskDb.Error != nil {
		return "error delete task", deleteTaskDb.Error
	} else if deleteTaskDb.RowsAffected < 1 {
		return "error delete task", errors.New("Data not found")
	}

	return "success delete task", nil
}
