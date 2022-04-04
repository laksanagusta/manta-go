package product

import "novocaine-dev/user"

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProductInput struct {
	Name          string `json:"name" binding:"required"`
	Serial_number string `json:"serial_number" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	Image_url     string `json:"image_url"`
	User_id       user.User
	Merchant_id   int `json:"merchant_id" binding:"required"`
}
