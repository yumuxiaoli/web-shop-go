package main

import (
	"fmt"
	"user/domain/repository"
	sr2 "user/domain/service"
	"user/handler"
	pb "user/proto/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)

	srv.Init()
	// 创建数据库
	db, err := gorm.Open("mysql",
		"root:123456@tcp(8.130.66.151:16161)/web_shop_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)

	// 只执行一次
	rp := repository.NewUserRepository(db)
	rp.InitTable()

	UserDataService := sr2.NewUserDataService(repository.NewUserRepository(db))
	// Register handler
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: UserDataService})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
