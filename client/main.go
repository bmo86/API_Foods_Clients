package main

import (
	"foods_API_GRPC/client/handlers"
	"foods_API_GRPC/server"

	"github.com/gin-gonic/gin"
)

func BindRoute(s server.Server, r *gin.Engine) {
	r.GET("/home", handlers.HandlerHome(s))
	r.POST("/food", handlers.HandlerCretedFood(s))

	r.GET("ws", handlers.HandlerWS(s))
}

func main() {

}
