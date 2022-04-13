package transactionProduct

import (
	"time"
)

type TransactionProductFormatter struct {
	ID             int       `json:"id"`
	Product_id     int       `json:"productId"`
	Transaction_Id int       `json:"transactionId"`
	ProductName    string    `json:"productName"`
	Qty            int       `json:"qty"`
	UnitPrice      int       `json:"unitPrice"`
	SubTotal       int       `json:"subTotal"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func FormatTransactionProduct(transactionProduct TransactionProduct) TransactionProductFormatter {
	transactionProductFormatter := TransactionProductFormatter{}
	transactionProductFormatter.ID = transactionProduct.ID
	transactionProductFormatter.Product_id = transactionProduct.Product_id
	transactionProductFormatter.Transaction_Id = transactionProduct.Transaction_Id
	transactionProductFormatter.ProductName = transactionProduct.ProductName
	transactionProductFormatter.Qty = transactionProduct.Qty
	transactionProductFormatter.UnitPrice = transactionProduct.UnitPrice
	transactionProductFormatter.SubTotal = transactionProduct.SubTotal
	transactionProductFormatter.CreatedAt = transactionProduct.CreatedAt
	transactionProductFormatter.UpdatedAt = transactionProduct.UpdatedAt

	return transactionProductFormatter
}

func FormatTasks(transactionProducts []TransactionProduct) []TransactionProductFormatter {
	transactionProductsFormatter := []TransactionProductFormatter{}

	for _, task := range transactionProducts {
		taskFormatter := FormatTransactionProduct(task)
		transactionProductsFormatter = append(transactionProductsFormatter, taskFormatter)
	}

	return transactionProductsFormatter
}
