package product

type ProductFormatter struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Serial_number string `json:"serial_number"`
	Price         int    `json:"price"`
	Image_url     string `json:"image_url"`
	Merchant_id   int    `json:"merchant_id"`
}

func FormatProduct(product Product) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.Name = product.Name
	productFormatter.Serial_number = product.Serial_number
	productFormatter.Price = product.Price
	productFormatter.Image_url = product.Image_url
	productFormatter.Merchant_id = product.Merchant_id

	return productFormatter
}

func FormatProducts(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
