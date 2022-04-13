package transaction

type TransactionInput struct {
	ID                int    `json:"id"`
	TransactionTitle  string `json:"transactionTitle"`
	TransactionStatus string `json:"transactionStatus"`
	UserID            int    `json:"userId"`
}

type AddToCartInput struct {
	Product_id     int    `json:"productId" binding:"required"`
	Transaction_Id int    `json:"transactionId"`
	User_id        int    `json:"userId"`
	ProductName    string `json:"productName"`
	UnitPrice      int    `json:"unitPrice" validate:"required,min=1"`
	Qty            int    `json:"qty" validate:"required,min=1"`
}

type FindById struct {
	ID int `uri:"id"`
}

type FindByUserId struct {
	UserID int `uri:"userId"`
}
