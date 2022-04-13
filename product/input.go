package product

import "novocaine-dev/user"

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProductInput struct {
	Name           string `json:"name" binding:"required"`
	Serial_number  string `json:"serial_number" binding:"required"`
	Price          int    `json:"price" binding:"required"`
	Image_url      string `json:"image_url"`
	User_id        user.User
	Organizaton_id int `json:"organization_id" binding:"required"`
}
