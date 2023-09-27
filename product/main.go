package main

import (
	"fmt"
	"product/common"
	"product/handler"
	pb "product/proto/product"

	"product/domain/repository"
	sr2 "product/domain/service"

	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("8.130.66.151", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistery := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"8.130.66.151:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("product", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	fmt.Println(mysqlInfo)
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+mysqlInfo.Port+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// 禁止复表
	db.SingularTable(true)

	repository.NewProductRepository(db).InitTable()

	ProductDataService := sr2.NewProductDataService(repository.NewProductRepository(db))

	// Create service
	srv := micro.NewService(
		micro.Name("product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		// 添加注册中心
		micro.Registry(consulRegistery),
		// 绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	srv.Init()
	// Register handler
	pb.RegisterProductHandler(srv.Server(), &handler.Product{ProductDataService: ProductDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
