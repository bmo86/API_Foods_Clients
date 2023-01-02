package handlers

import (
	"foods_API_GRPC/models"
	"foods_API_GRPC/proto/foodspb"
	"foods_API_GRPC/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerHome(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to my API with golang, grpc, jwt, gin and gorm/postgres",
		})
	}
}

func HandlerWS(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s.Hub().HandlerWs(ctx.Writer, ctx.Request)
	}
}

func HandlerCretedFood(s server.Server, food foodspb.FoodServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var f models.Food

		if err := ctx.ShouldBind(&f); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Data incorrect",
			})
			return
		}

		data := foodspb.CreatedFoodRequest{
			Name:        f.Name,
			Price:       float32(f.Price),
			Ingredients: f.Ingredients,
		}

		//method grpc add
		food.CreatedFood(ctx, &data)

		s.Hub().Broadcast(&data, nil)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "created",
		})
	}
}
