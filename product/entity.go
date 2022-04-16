package product

import "time"

type Product struct {
	ID             int
	Name           string
	Serial_number  string
	Price          int
	Image_url      string
	User_id        int
	OrganizationId int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BulkProducts struct {
	Name         string
	Price        int
	SerialNumber string
}
