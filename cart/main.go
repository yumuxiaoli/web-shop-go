package main

import (
	"cart/domain/repository"
	"cart/domain/service"
	"cart/handler"
	pb "cart/proto"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	sr2 "github.com/micro/go-micro/v2/registry/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/yumuxiaoli/common"
)

var QPS = 100

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("8.130.66.151", 8500, "micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consul := sr2.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"8.130.66.151:8500",
		}
	})
	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 数据库连接
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+mysqlInfo.Port+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// 禁止副表
	db.SingularTable(true)

	// 第一次初始化
	err = repository.NewCartRepository(db).InitTable()
	if err != nil {
		log.Error(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Name("cart"),
		micro.Version("latest"),
		// 暴露地址
		micro.Address("0.0.0.0:8087"),
		micro.Registry(consul),
		// 链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	srv.Init()

	cartDataService := service.NewCartDataService(repository.NewCartRepository(db))

	// Register handler
	pb.RegisterCartHandler(srv.Server(), &handler.Cart{CartDataService: cartDataService})

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
