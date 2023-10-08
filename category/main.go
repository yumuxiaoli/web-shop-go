package main

import (
	"category/common"
	"category/handler"
	pb "category/proto"
	"fmt"

	"category/domain/repository"
	sr2 "category/domain/service"

	log "github.com/micro/go-micro/v2/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("8.130.66.151", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistery := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"8.130.66.151:8500",
		}
	})
	// Create service
	srv := micro.NewService(
		micro.Name("category"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul 作为注册中心
		micro.Registry(consulRegistery),
	)
	// 获取mysql的配置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	fmt.Println(mysqlInfo)
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+mysqlInfo.Port+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// 禁止复表
	db.SingularTable(true)

	// rp := repository.NewCategoryRepository(db)
	// rp.InitTable()

	srv.Init()

	categoryDataService := sr2.NewCategoryDataService(repository.NewCategoryRepository(db))

	err = pb.RegisterCategoryHandler(srv.Server(), &handler.Category{CategoryDateService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
