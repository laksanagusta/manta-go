package transaction

import (
	"novocaine-dev/transactionProduct"
	"novocaine-dev/user"
	"time"
)

type Transaction struct {
	ID                  int
	TransactionTitle    string
	TransactionStatus   string
	User_id             int
	User                user.User
	TransactionProducts []transactionProduct.TransactionProduct
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
