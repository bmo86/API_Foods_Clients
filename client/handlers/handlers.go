package handlers

import (
	"foods_API_GRPC/models"
	"foods_API_GRPC/proto/foodspb"
	"foods_API_GRPC/server"
	"net/http"
	"strconv"

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

func HandlerCretedFood(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var f models.FoodWithIngredients

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
		res, err := s.Proto().FoodServiceClient.CreatedFood(ctx, &data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		s.Hub().Broadcast(res, nil)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "created",
			"data":    res,
		})
	}
}

func HandlerGetFoood(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idReq := ctx.Param("id")
		id, err := strconv.ParseInt(idReq, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "id - not found",
			})
			return
		}

		res, err := s.Proto().FoodServiceClient.GetFood(ctx, &foodspb.FoodRequest{Id: id})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "get - food",
			"data":    res,
		})
	}
}
