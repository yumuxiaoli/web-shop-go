# web-shop-go
微服务商城项目

代码结构
* user
```
.
├── Dockerfile
├── domain
│   ├── model
│   │   └── user.go
│   ├── repository
│   │   └── user_repository.go
│   └── service
│       └── user_data_service.go
├── go.mod
├── go.sum
├── handler
│   └── user.go
├── main.go
├── Makefile
├── proto
│   ├── user.pb.go
│   ├── user.pb.micro.go
│   └── user.proto
├── README.md
└── user
```

* product
```
.
├── common
│   ├── config.go
│   ├── jaeger.go
│   ├── mysql.go
│   └── swap.go
├── domain
│   ├── model
│   │   ├── product.go
│   │   ├── product_image.go
│   │   ├── product_seo.go
│   │   └── product_size.go
│   ├── repository
│   │   └── product_repository.go
│   └── service
│       └── product_data.go
├── go.mod
├── go.sum
├── handler
│   └── product.go
├── main.go
├── Makefile
├── productClient.go
├── proto
│   ├── product.pb.go
│   ├── product.pb.micro.go
│   └── product.proto
└── README.md
```
* category
```
.
├── common
│   ├── config.go
│   ├── mysql.go
│   └── swap.go
├── Dockerfile
├── domain
│   ├── model
│   │   └── category.go
│   ├── repository
│   │   └── category_repository.go
│   └── service
│       └── category_date.go
├── go.mod
├── go.sum
├── handler
│   └── category.go
├── main.go
├── Makefile
├── proto
│   ├── category.pb.go
│   ├── category.pb.micro.go
│   └── category.proto
└── README.md
```
