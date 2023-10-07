package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdataCart(Cart *model.Cart) (err error)
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

func NewCartDataService(CartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{CartRepository: CartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func (u *CartDataService) AddCart(Cart *model.Cart) (int64, error) {
	return u.CartRepository.CreateCart(Cart)
}

func (u *CartDataService) DeleteCart(CartID int64) error {
	return u.CartRepository.DeleteCartByID(CartID)
}

func (u *CartDataService) UpdataCart(Cart *model.Cart) (err error) {
	return u.CartRepository.UpdataCart(Cart)
}

func (u *CartDataService) FindCartByID(CartID int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(CartID)
}

func (u *CartDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userID)
}

func (u *CartDataService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}

func (u *CartDataService) IncrNum(cartID int64, num int64) error {
	return u.CartRepository.IncrNum(cartID, num)
}

func (u *CartDataService) DecrNum(cartID int64, num int64) error {
	return u.CartRepository.DecrNum(cartID, num)
}
