package transactionProduct

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(transactionProduct TransactionProduct) (TransactionProduct, error)
	Delete(transactionProductId int) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transactionProduct TransactionProduct) (TransactionProduct, error) {
	err := r.db.Create(&transactionProduct).Error
	if err != nil {
		return transactionProduct, err
	}
	return transactionProduct, nil
}

func (r *repository) Delete(transactionProductId int) (string, error) {
	err := r.db.Delete(&TransactionProduct{}, transactionProductId).Error
	if err != nil {
		return "error delete items", err
	}

	return "success delete items", nil
}
