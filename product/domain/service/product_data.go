package service

import (
	"product/domain/model"
	"product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdataProduct(Product *model.Product) (err error)
	FindAllProduct() ([]model.Product, error)
	FindProductByID(int64) (*model.Product, error)
}

func NewProductDataService(ProductRepository repository.IProductRepository) IProductDataService {
	return &ProductDataService{ProductRepository: ProductRepository}
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

func (u *ProductDataService) AddProduct(Product *model.Product) (int64, error) {
	return u.ProductRepository.CreateProduct(Product)
}

func (u *ProductDataService) DeleteProduct(ProductID int64) error {
	return u.ProductRepository.DeleteProductByID(ProductID)
}

func (u *ProductDataService) UpdataProduct(Product *model.Product) (err error) {
	return u.ProductRepository.UpdataProduct(Product)
}

func (u *ProductDataService) FindProductByID(ProductID int64) (*model.Product, error) {
	return u.ProductRepository.FindProductByID(ProductID)
}

func (u *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return u.ProductRepository.FindAll()
}
