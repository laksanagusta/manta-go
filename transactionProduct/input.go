package transactionProduct

type TransactionProductInput struct {
	Product_id     int    `json:"productId" binding:"required"`
	Transaction_Id int    `json:"transactionId" binding:"required"`
	ProductName    string `json:"productName" binding:"required"`
	UnitPrice      int    `json:"unitPrice" binding:"required"`
	SubTotal       int    `json:"subTotal" binding:"required"`
	Qty            int    `json:"qty" binding:"required"`
	UserID         int    `json:"userId"`
}
