package repository

import (
	"product/domain/model"

	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdataProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

func (u *ProductRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Product{}, &model.ProductSeo{}, &model.ProductImage{}, &model.ProductSize{}).Error
}

func (u *ProductRepository) FindProductByID(ProductID int64) (Product *model.Product, err error) {
	Product = &model.Product{}
	return Product, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(Product, ProductID).Error
}

func (u *ProductRepository) CreateProduct(Product *model.Product) (ProductID int64, err error) {
	return Product.ID, u.mysqlDb.Create(Product).Error
}

func (u *ProductRepository) DeleteProductByID(ProductID int64) error {
	// 开启事务
	tx := u.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 删除
	if err := tx.Unscoped().Where("id = ?", ProductID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("images_product_id = ?", ProductID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", ProductID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", ProductID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *ProductRepository) UpdataProduct(Product *model.Product) error {
	return u.mysqlDb.Model(Product).Update(&Product).Error
}

func (u *ProductRepository) FindAll() (Productlist []model.Product, err error) {
	return Productlist, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&Productlist).Error
}
