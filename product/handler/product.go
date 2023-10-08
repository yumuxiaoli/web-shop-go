package handler

import (
	"context"
	"product/common"
	"product/domain/model"
	"product/domain/service"
	product "product/proto"
)

const (
	UpdateSuccess = "更新成功"
	DeleteSuccess = "删除成功"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p *Product) AddProduct(ctx context.Context, prq *product.ProductInfo, prp *product.ResponseProduct) error {
	product := &model.Product{}
	if err := common.SwapTo(prq, product); err != nil {
		return err
	}
	productID, err := p.ProductDataService.AddProduct(product)
	if err != nil {
		return err
	}
	prp.ProductId = productID
	return nil
}

// 根据ID查找商品
func (p *Product) FindProductById(ctx context.Context, prq *product.RequestID, prp *product.ProductInfo) error {
	productData, err := p.ProductDataService.FindProductByID(prq.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData, prp); err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, prq *product.ProductInfo, prp *product.Response) error {
	product := &model.Product{}
	if err := common.SwapTo(prq, product); err != nil {
		return err
	}
	prp.Msg = UpdateSuccess
	return nil
}

func (p *Product) DeleteProductByID(ctx context.Context, prq *product.RequestID, prp *product.Response) error {
	if err := p.ProductDataService.DeleteProduct(prq.ProductId); err != nil {
		return err
	}
	prp.Msg = DeleteSuccess
	return nil
}

func (p *Product) FindAllProduct(ctx context.Context, prq *product.RequestAll, prp *product.AllProduct) error {
	productAll, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}

	for _, v := range productAll {
		productInfo := &product.ProductInfo{}
		err := common.SwapTo(v, productInfo)
		if err != nil {
			return err
		}
		prp.ProductInfo = append(prp.ProductInfo, productInfo)
	}
	return nil
}
