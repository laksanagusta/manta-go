package product

import "time"

type Product struct {
	ID             int
	Name           string
	Serial_number  string
	Price          int
	Image_url      string
	User_id        int
	Organizaton_id int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
