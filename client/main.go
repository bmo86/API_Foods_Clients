package main

import (
	"context"
	"flag"
	"foods_API_GRPC/client/handlers"
	"foods_API_GRPC/proto/foodspb"
	"foods_API_GRPC/server"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:5060", "the address to connect to")
)

func BindRoute(s server.Server, r *gin.Engine) {
	r.GET("/home", handlers.HandlerHome(s))
	//foods
	r.POST("/food", handlers.HandlerCretedFood(s))
	r.GET("/food/:id", handlers.HandlerGetFoood(s))
	r.GET("/foods/:page", handlers.HandlerGetFoods(s))
	r.GET("/ws", handlers.HandlerWS(s))
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	client := foodspb.NewFoodServiceClient(conn)
	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret: "xd",
		Port:      "8080",
		UrlDB:     "xd",
	}, &server.Proto{FoodServiceClient: client})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoute)
}
