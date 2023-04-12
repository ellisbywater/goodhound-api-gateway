package auth

import (
	"github.com/ellisbywater/grpc-api-gateway/pkg/auth/routes"
	"github.com/ellisbywater/grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	return svc
}

func (c *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, c.Client)
}

func (c *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, c.Client)
}
