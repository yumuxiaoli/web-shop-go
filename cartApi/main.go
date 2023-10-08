package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/opentracing/opentracing-go"

	"github.com/yumuxiaoli/common"
	cart "github.com/yumuxiaoli/web-shop-go/cart/proto"
	"github.com/yumuxiaoli/web-shop-go/cartApi/handler"
	pb "github.com/yumuxiaoli/web-shop-go/cartApi/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	col2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

func main() {
	// 注册中心
	consul := col2.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"8.130.66.151:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.api.cartApi", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	// 启动端口
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// Create service
	srv := micro.NewService(
		micro.Name("cartapi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		// 增加 consul 注册中心
		micro.Registry(consul),
		// 添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 增加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		// 添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	srv.Init()

	cartService := cart.NewCartService("go.micro.service.cart", srv.Client())
	// Register handler
	if err := pb.RegisterCartApiHandler(srv.Server(), &handler.CartApi{CartService: cartService}); err != nil {
		log.Error(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rep interface{}, opts ...client.CallOption) error {
	return hystrix.Do(
		req.Service()+"."+req.Endpoint(), func() error {
			// run 正常进行逻辑
			fmt.Println(req.Service() + "." + req.Endpoint())
			return c.Client.Call(ctx, req, rep, opts...)
		}, func(err error) error {
			fmt.Println(err)
			return nil
		})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{}
	}
}
