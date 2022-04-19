package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(transaction Transaction) (Transaction, error)
	FindOngoingTransaction(userId int) (Transaction, error)
	AddTransactionProduct(transaction Transaction) (Transaction, error)
	FindById(transactionId int) (Transaction, error)
	FindTransactionByUser(userId int) ([]Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindOngoingTransaction(userId int) (Transaction, error) {
	transaction := Transaction{}
	err := r.db.Preload("TransactionProducts").Where("user_id = ?", userId).Last(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) AddTransactionProduct(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindById(transactionId int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Preload("TransactionProducts").Preload("User").Where("id = ?", transactionId).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindTransactionByUser(userId int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Preload("TransactionProducts").Preload("User").Where("user_id = ?", userId).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
