package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/micro/go-micro/v2/util/log"
	cart "github.com/yumuxiaoli/web-shop-go/cart/proto"
	cartApi "github.com/yumuxiaoli/web-shop-go/cartApi/proto"
)

type CartApi struct{ CartService cart.CartService }

func (e *CartApi) FindAll(ctx context.Context, Arq *cartApi.Request, Arp *cartApi.Response) error {
	log.Info("接受到 /cartApi/FindAll 访问请求")
	if _, ok := Arq.Get["user_id"]; !ok {
		Arp.StatusCode = 500
		return errors.New("参数异常")
	}
	userIdString := Arq.Get["user_id"].Values[0]
	fmt.Println(userIdString)
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return err
	}
	// 获取购物车所有商品
	cartAll, err := e.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	// 数据类型转化
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	Arp.StatusCode = 200
	Arp.Body = string(b)
	return nil
}
