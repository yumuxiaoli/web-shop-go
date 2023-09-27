package handler

import (
	"category/common"
	"category/domain/model"
	"category/domain/service"
	"category/proto/category"
	"context"

	"github.com/micro/go-micro/v2/util/log"
)

const (
	CategorySucess       = "分类成功"
	CategoryUpdateSucess = "分类更新成功"
	CategoryDeleteSucess = "分类删除成功"
)

type Category struct {
	CategoryDateService service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, crq *category.CategoryRequest, crp *category.CreateCategoryResponse) error {
	category := &model.Category{}
	// 赋值
	err := common.SwapTo(crq, category)
	if err != nil {
		return err
	}
	categoryID, err := c.CategoryDateService.AddCategory(category)
	if err != nil {
		return err
	}
	crp.Message = CategorySucess
	crp.CategoryId = categoryID
	return nil
}
func (c *Category) UpdateCategory(ctx context.Context, crq *category.CategoryRequest, crp *category.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(crq, category)
	if err != nil {
		return err
	}
	err = c.CategoryDateService.UpdataCategory(category)
	if err != nil {
		return err
	}
	crp.Message = CategoryUpdateSucess
	return nil
}
func (c *Category) DeleteCategory(ctx context.Context, crq *category.DeleteCategoryRequest, crp *category.DeleteCategoryResponse) error {
	err := c.CategoryDateService.DeleteCategory(crq.CategoryId)
	if err != nil {
		return err
	}
	crp.Message = CategoryDeleteSucess
	return nil
}
func (c *Category) FindCategoryByName(ctx context.Context, crq *category.FindByNameRequset, crp *category.CategoryResponse) error {
	category, err := c.CategoryDateService.FindCategoryByName(crq.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, crp)
}
func (c *Category) FindCategoryByID(ctx context.Context, crq *category.FindByIDRequest, crp *category.CategoryResponse) error {
	category, err := c.CategoryDateService.FindCategoryByID(crq.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, crp)
}
func (c *Category) FindCategoryByLevel(ctx context.Context, crq *category.FindByLevelRequest, crp *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDateService.FindCategoryLevel(crq.CategoryLevel)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, crp)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, crq *category.FindByParentRequest, crp *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDateService.FindCategoryByParent(crq.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, crp)
	return nil
}
func (c *Category) FindAllCategory(ctx context.Context, crq *category.FindAllRequest, crp *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDateService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, crp)
	return nil
}

func categoryToResponse(categorySlic []model.Category, crp *category.FindAllResponse) {
	for _, cg := range categorySlic {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		crp.Category = append(crp.Category, cr)
	}
}
