package routes

import (
	"context"
	"github.com/ellisbywater/grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
