package product

import (
	"github.com/ellisbywater/grpc-api-gateway/pkg/auth"
	"github.com/ellisbywater/grpc-api-gateway/pkg/config"
	"github.com/ellisbywater/grpc-api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.GET("/:id", svc.FindOne)
	routes.POST("/create", svc.CreateProduct)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
