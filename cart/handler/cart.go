package handler

import (
	"context"

	"github.com/yumuxiaoli/common"
	"github.com/yumuxiaoli/web-shop-go/cart/domain/model"
	"github.com/yumuxiaoli/web-shop-go/cart/domain/service"
	cart "github.com/yumuxiaoli/web-shop-go/cart/proto"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (h *Cart) AddCart(ctx context.Context, crq *cart.CartInfo, crp *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(crq, cart)
	crp.CartId, err = h.CartDataService.AddCart(cart)
	return err
}

func (h *Cart) CleanCart(ctx context.Context, crq *cart.Clean, crp *cart.Response) error {
	if err := h.CartDataService.CleanCart(crq.UserId); err != nil {
		return err
	}
	crp.Msg = "购物车清空成功"
	return nil
}

func (h *Cart) Incr(ctx context.Context, crq *cart.Item, crp *cart.Response) error {
	if err := h.CartDataService.IncrNum(crq.Id, crq.ChangeNum); err != nil {
		return err
	}
	crp.Msg = "购物车添加成功"
	return nil
}

func (h *Cart) Decr(ctx context.Context, crq *cart.Item, crp *cart.Response) error {
	if err := h.CartDataService.DecrNum(crq.Id, crq.ChangeNum); err != nil {
		return err
	}
	crp.Msg = "购物车减少成功"
	return nil
}
func (h *Cart) DeleteItemByID(ctx context.Context, crq *cart.CartID, crp *cart.Response) error {
	if err := h.CartDataService.DeleteCart(crq.Id); err != nil {
		return err
	}
	crp.Msg = "购物车删除成功"
	return nil
}

func (h *Cart) GetAll(ctx context.Context, crq *cart.CartFindAll, crp *cart.CartAll) error {
	cartAll, err := h.CartDataService.FindAllCart(crq.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		crp.CartInfo = append(crp.CartInfo, cart)
	}
	return nil
}
