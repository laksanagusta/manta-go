package transactionProduct

type Service interface {
	CreateTransactionProduct(input TransactionProductInput) (TransactionProduct, error)
	Delete(transactionProductId int) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransactionProduct(input TransactionProductInput) (TransactionProduct, error) {
	transactionProduct := TransactionProduct{}
	transactionProduct.Transaction_Id = input.Transaction_Id
	transactionProduct.Product_id = input.Product_id
	transactionProduct.ProductName = input.ProductName
	transactionProduct.UnitPrice = input.UnitPrice
	transactionProduct.SubTotal = input.SubTotal

	newTransaction, err := s.repository.Save(transactionProduct)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

func (s *service) Delete(transactionProductId int) (string, error) {
	deleteTransactionProduct, err := s.repository.Delete(transactionProductId)
	if err != nil {
		return deleteTransactionProduct, err
	}
	return deleteTransactionProduct, nil
}
