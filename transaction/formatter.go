package transaction

import "time"

type TransactionFormatter struct {
	ID                 int                           `json:"id"`
	TransactionTitle   string                        `json:"transactionTitle"`
	TransactionStatus  string                        `json:"transactionStatus"`
	User               TransactionUserFormatter      `json:"user"`
	TransactionDetails []TransactionDetailsFormatter `json:"transactionDetails"`
	CreatedAt          time.Time                     `json:"createdAt"`
	UpdatedAt          time.Time                     `json:"updatedAt"`
}

type TransactionDetailsFormatter struct {
	ID          int       `json:"id"`
	ProductName string    `json:"productName"`
	Qty         int       `json:"qty"`
	SubTotal    int       `json:"subTotal"`
	UnitPrice   int       `json:"unitPrice"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TransactionUserFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	transactionFormatter := TransactionFormatter{}
	transactionFormatter.ID = transaction.ID
	transactionFormatter.TransactionTitle = transaction.TransactionTitle
	transactionFormatter.TransactionStatus = transaction.TransactionStatus

	transactionUserFormatter := TransactionUserFormatter{}
	transactionUserFormatter.Name = transaction.User.Name
	transactionUserFormatter.ID = transaction.User_id
	transactionUserFormatter.Username = transaction.User.Username

	transactionFormatter.User = transactionUserFormatter

	transactionDetails := []TransactionDetailsFormatter{}
	for _, v := range transaction.TransactionProducts {
		transactionDetailsOut := TransactionDetailsFormatter{}
		transactionDetailsOut.ID = v.ID
		transactionDetailsOut.ProductName = v.ProductName
		transactionDetailsOut.Qty = v.Qty
		transactionDetailsOut.SubTotal = v.SubTotal
		transactionDetailsOut.UnitPrice = v.UnitPrice
		transactionDetailsOut.UpdatedAt = v.UpdatedAt
		transactionDetailsOut.CreatedAt = v.CreatedAt
		transactionDetails = append(transactionDetails, transactionDetailsOut)
	}
	transactionFormatter.TransactionDetails = transactionDetails
	transactionFormatter.CreatedAt = transaction.CreatedAt
	transactionFormatter.UpdatedAt = transaction.UpdatedAt

	return transactionFormatter
}

func FormatTransactions(transaction []Transaction) []TransactionFormatter {
	transactionsFormatter := []TransactionFormatter{}

	for _, v := range transaction {
		transactionFormatter := FormatTransaction(v)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	return transactionsFormatter
}
