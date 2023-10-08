package main

import (
	"context"
	"fmt"
	"product/common"
	product "product/proto"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	// 注册中心
	consulRegistery := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"8.130.66.151:8500",
		}
	})
	// 链路追踪
	t, io, err := common.NewTracer("product.client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Create service
	srv := micro.NewService(
		micro.Name("product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8084"),
		// 添加注册中心
		micro.Registry(consulRegistery),
		// 绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)
	productService := product.NewProductService("product", srv.Client())
	productAdd := &product.ProductInfo{
		ProductName:        "yuyu",
		ProductSku:         "yumu",
		ProductPrice:       1.4,
		ProductDescription: "yumu",
		ProductImage: []*product.ProductImage{
			{
				ImagesName: "xiaoyu",
				ImagesCode: "xiaoyu01",
				ImagesUrl:  "xioayu01",
			},
			{
				ImagesName: "xiaoyu2",
				ImagesCode: "xiaoyu02",
				ImagesUrl:  "xioayu02",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "xiaoyu",
				SizeCode: "xioayu-code",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "xiaoyu-seo",
			SeoKeywords:    "xiaoyu-seo",
			SeoDescription: "xiaoyu-seo",
			SeoCode:        "xioayu-code",
		},
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
