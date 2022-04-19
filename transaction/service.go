package transaction

import (
	"errors"
	"novocaine-dev/transactionProduct"
)

type Service interface {
	CreateTransaction(input TransactionInput) (Transaction, error)
	FindOngoingTransaction(userId int) (Transaction, error)
	FindById(transactionId int) (Transaction, error)
	AddToCart(input AddToCartInput) (Transaction, error)
	FindTransactionByUser(userId int) ([]Transaction, error)
	UpdateTransaction(id FindById, inputData TransactionInput) (Transaction, error)
}

type service struct {
	repository             Repository
	transactionProductRepo transactionProduct.Repository
}

func NewService(repository Repository, transactionProductRepo transactionProduct.Repository) *service {
	return &service{repository, transactionProductRepo}
}

func (s *service) CreateTransaction(input TransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.TransactionStatus = "ongoing"
	transaction.TransactionTitle = input.TransactionTitle
	transaction.User_id = input.UserID

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

func (s *service) FindOngoingTransaction(userId int) (Transaction, error) {

	newTransaction, err := s.repository.FindOngoingTransaction(userId)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

func (s *service) AddToCart(input AddToCartInput) (Transaction, error) {
	transactionProducts := transactionProduct.TransactionProduct{}
	transactionProducts.ProductName = input.ProductName
	transactionProducts.Product_id = input.Product_id
	transactionProducts.ProductName = input.ProductName
	transactionProducts.SubTotal = input.UnitPrice * input.Qty
	transactionProducts.UnitPrice = input.UnitPrice
	transactionProducts.Qty = input.Qty

	ongoingTransaction, err := s.repository.FindOngoingTransaction(input.User_id)
	transactionProducts.Transaction_Id = ongoingTransaction.ID
	if ongoingTransaction.ID == 0 {
		transaction := Transaction{}
		transaction.TransactionTitle = "Transaction - Title"
		transaction.TransactionStatus = "ongoing"
		transaction.User_id = input.User_id
		ongoingTransaction, err = s.repository.Save(transaction)

		if err != nil {
			return ongoingTransaction, err
		}
		transactionProducts.Transaction_Id = ongoingTransaction.ID
	}

	transactionDetails, err := s.transactionProductRepo.Save(transactionProducts)
	transactionProducts.ID = transactionDetails.ID
	transactionProducts.CreatedAt = transactionDetails.CreatedAt
	transactionProducts.UpdatedAt = transactionDetails.UpdatedAt

	ongoingTransaction.TransactionProducts = append(ongoingTransaction.TransactionProducts, transactionProducts)

	if err != nil {
		return ongoingTransaction, err
	}
	return ongoingTransaction, nil
}

func (s *service) FindById(transactionId int) (Transaction, error) {
	transaction, err := s.repository.FindById(transactionId)

	if transaction.ID == 0 {
		return transaction, errors.New("Transaction not found")
	}

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) FindTransactionByUser(userId int) ([]Transaction, error) {
	transaction, err := s.repository.FindTransactionByUser(userId)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) UpdateTransaction(id FindById, inputData TransactionInput) (Transaction, error) {
	transaction, err := s.repository.FindById(id.ID)
	if err != nil {
		return transaction, err
	}

	transaction.TransactionTitle = inputData.TransactionTitle
	transaction.TransactionStatus = inputData.TransactionStatus

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return updatedTransaction, err
	}

	return updatedTransaction, nil

}
