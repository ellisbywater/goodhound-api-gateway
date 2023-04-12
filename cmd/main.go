package main

import (
	"github.com/ellisbywater/grpc-api-gateway/pkg/auth"
	"github.com/ellisbywater/grpc-api-gateway/pkg/config"
	"github.com/ellisbywater/grpc-api-gateway/pkg/order"
	"github.com/ellisbywater/grpc-api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}
