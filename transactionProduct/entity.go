package transactionProduct

import "time"

type TransactionProduct struct {
	ID             int
	Product_id     int
	Transaction_Id int
	ProductName    string
	Qty            int
	UnitPrice      int
	SubTotal       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ID struct {
	ID int `uri:"id"`
}
