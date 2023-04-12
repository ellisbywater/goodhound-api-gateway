package routes

import (
	"context"
	"github.com/ellisbywater/grpc-api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId, _ := ctx.Get("userId")
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
