package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdataCategory(Category *model.Category) (err error)
	FindCategoryByName(string) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByID(int64) (*model.Category, error)
	FindCategoryLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

func NewCategoryDataService(CategoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{CategoryRepository: CategoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

func (u *CategoryDataService) AddCategory(Category *model.Category) (int64, error) {
	return u.CategoryRepository.CreateCategory(Category)
}

func (u *CategoryDataService) DeleteCategory(CategoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(CategoryID)
}

func (u *CategoryDataService) UpdataCategory(Category *model.Category) (err error) {
	return u.CategoryRepository.UpdataCategory(Category)
}

func (u *CategoryDataService) FindCategoryByName(CategoryName string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(CategoryName)
}

func (u *CategoryDataService) FindCategoryByID(CategoryID int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(CategoryID)
}

func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

func (u *CategoryDataService) FindCategoryLevel(level uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryLevel(level)
}

func (u *CategoryDataService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}
