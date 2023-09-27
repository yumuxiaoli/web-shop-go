package repository

import (
	"category/domain/model"

	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdataCategory(*model.Category) error
	FindAll() ([]model.Category, error)
	FindCategoryLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

func (u *CategoryRepository) FindCategoryByName(name string) (Category *model.Category, err error) {
	Category = &model.Category{}
	return Category, u.mysqlDb.Where("category_name = ?", name).Find(Category).Error
}

func (u *CategoryRepository) FindCategoryByID(CategoryID int64) (Category *model.Category, err error) {
	Category = &model.Category{}
	return Category, u.mysqlDb.First(Category, CategoryID).Error
}

func (u *CategoryRepository) CreateCategory(Category *model.Category) (CategoryID int64, err error) {
	return Category.ID, u.mysqlDb.Create(Category).Error
}

func (u *CategoryRepository) DeleteCategoryByID(CategoryID int64) error {
	return u.mysqlDb.Where("id = ?", CategoryID).Delete(&model.Category{}).Error
}

func (u *CategoryRepository) UpdataCategory(Category *model.Category) error {
	return u.mysqlDb.Model(Category).Update(&Category).Error
}

func (u *CategoryRepository) FindAll() (Categorylist []model.Category, err error) {
	return Categorylist, u.mysqlDb.Find(&Categorylist).Error
}

func (u *CategoryRepository) FindCategoryLevel(level uint32) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("category_level = ?", level).Find(categorySlice).Error
}
func (u *CategoryRepository) FindCategoryByParent(parent int64) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("category_parent = ?", parent).Find(categorySlice).Error
}
