package product

import (
	"fmt"
	"strconv"
)

type Service interface {
	GetProducts(id int) ([]Product, error)
	GetProductById(id int) (Product, error)
	CreateProduct(input CreateProductInput) (Product, error)
	UpdateProduct(inputID GetProductDetailInput, inputData CreateProductInput) (Product, error)
	DeleteProduct(inputID GetProductDetailInput) (string, error)
	SaveImage(ID int, fileLocation string) (Product, error)
	CreateProductBulk(workerIndex int, counter int, jobs []interface{}) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(id int) ([]Product, error) {
	if id != 0 {
		products, err := s.repository.FindByProductId(id)
		if err != nil {
			return products, err
		}

		return products, nil
	}

	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetProductById(id int) (Product, error) {
	products, err := s.repository.FindById(id)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) CreateProduct(input CreateProductInput) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Serial_number = input.Serial_number
	product.Price = input.Price
	product.OrganizationId = input.User_id.OrganizationId
	product.User_id = input.User_id.ID
	product.Image_url = input.Image_url
	newProducts, err := s.repository.Save(product)
	if err != nil {
		return newProducts, err
	}

	return newProducts, nil

}

func (s *service) UpdateProduct(inputID GetProductDetailInput, inputData CreateProductInput) (Product, error) {
	product, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return product, err
	}

	product.Name = inputData.Name
	product.Serial_number = inputData.Serial_number
	product.Price = inputData.Price

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil

}

func (s *service) DeleteProduct(inputID GetProductDetailInput) (string, error) {
	updatedProduct, err := s.repository.Delete(inputID.ID)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil

}

func (s *service) SaveImage(ID int, fileLocation string) (Product, error) {
	products, err := s.repository.FindById(ID)
	if err != nil {
		return products, err
	}

	products.Image_url = fileLocation
	updatedProduct, err := s.repository.Update(products)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil

}

func (s *service) CreateProductBulk(workerIndex int, counter int, jobs []interface{}) (Product, error) {
	product := Product{}
	product.Name = fmt.Sprintf("%v", jobs[0])
	product.Serial_number = fmt.Sprintf("%v", jobs[1])

	price := fmt.Sprintf("%v", jobs[2])
	convertPrice, err := strconv.Atoi(price)
	if err != nil {
		return product, err
	}
	product.Price = convertPrice

	organizationId := fmt.Sprintf("%v", jobs[3])
	convertOrganizationId, err := strconv.Atoi(organizationId)
	if err != nil {
		return product, err
	}
	product.OrganizationId = convertOrganizationId

	userId := fmt.Sprintf("%v", jobs[4])
	convertUserId, err := strconv.Atoi(userId)
	if err != nil {
		return product, err
	}
	product.User_id = convertUserId

	newProducts, err := s.repository.Save(product)
	if err != nil {
		return newProducts, err
	}

	return newProducts, nil
}
