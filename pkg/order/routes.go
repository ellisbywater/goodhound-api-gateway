package order

import (
	"github.com/ellisbywater/grpc-api-gateway/pkg/auth"
	"github.com/ellisbywater/grpc-api-gateway/pkg/config"
	"github.com/ellisbywater/grpc-api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/create", svc.CreateOrder)
	return svc
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
