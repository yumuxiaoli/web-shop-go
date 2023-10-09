package main

import (
	"fmt"

	col2 "github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	"github.com/jinzhu/gorm"

	ratelimit "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	op2 "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/yumuxiaoli/common"
	"github.com/yumuxiaoli/web-shop-go/cart/domain/repository"
	"github.com/yumuxiaoli/web-shop-go/cart/domain/service"
	"github.com/yumuxiaoli/web-shop-go/cart/handler"
	pb "github.com/yumuxiaoli/web-shop-go/cart/proto"
)

var QPS = 100

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("8.130.66.151", 8500, "micro/config")
	if err != nil {
		logger.Error(err)
	}
	// 注册中心
	consul := col2.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"8.130.66.151:8500",
		}
	})
	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		logger.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 数据库连接
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+mysqlInfo.Port+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()
	// 禁止副表
	db.SingularTable(true)

	// 第一次初始化
	err = repository.NewCartRepository(db).InitTable()
	if err != nil {
		logger.Error(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Name("cart"),
		micro.Version("latest"),
		// 暴露地址
		micro.Address("0.0.0.0:8087"),
		micro.Registry(consul),
		// 链路追踪
		micro.WrapHandler(op2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	srv.Init()

	cartDataService := service.NewCartDataService(repository.NewCartRepository(db))

	// Register handler
	err = pb.RegisterCartHandler(srv.Server(), &handler.Cart{CartDataService: cartDataService})
	if err != nil {
		logger.Error(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
