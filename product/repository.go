package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByProductId(id int) ([]Product, error)
	FindById(id int) (Product, error)
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(productID int) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var product []Product
	err := r.db.Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByProductId(id int) ([]Product, error) {
	var product []Product
	err := r.db.Where("id = ?", id).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindById(id int) (Product, error) {
	var product Product
	err := r.db.Where("id = ?", id).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Delete(productID int) (string, error) {
	err := r.db.Delete(&Product{}, productID).Error
	if err != nil {
		return "error delete product", err
	}

	return "success delete product", nil
}
