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
				"message": "Data incorrect",
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

func HandlerGetFoods(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pageReq := ctx.Param("page")

		page, err := strconv.ParseInt(pageReq, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if page == 0 {
			page = 5
		}

		res, err := s.Proto().GetFoods(ctx, &foodspb.GetFoodsRequest{Page: page})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "get",
			"foods":   res,
		})

	}
}

func HandlerDeleteFoods(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idReq := ctx.Param("id")
		id, err := strconv.ParseInt(idReq, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		res, err := s.Proto().DeleteFood(ctx, &foodspb.FoodRequest{Id: id})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Delete",
			"food":    res,
		})
	}
}

func HandlerUpdateFoods(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var food models.FoodUpdate

		if err := ctx.ShouldBind(&food); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		idReq := ctx.Param("id")
		id, err := strconv.ParseInt(idReq, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		data := &foodspb.FoodUpdateRequest{
			Id:    id,
			Name:  food.Name,
			Price: float32(food.Price),
		}

		res, err := s.Proto().UpdateFood(ctx, data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "update",
			"update":  res,
		})

	}
}
