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
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	//defer conn.Close()

	client := foodspb.NewFoodServiceClient(conn)

	r.GET("/home", handlers.HandlerHome(s))
	r.POST("/food", handlers.HandlerCretedFood(s, client))
	r.GET("/food/:id", handlers.HandlerCretedFood(s, client))

	r.GET("/ws", handlers.HandlerWS(s))
}

func main() {
	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret: "xd",
		Port:      "8080",
		UrlDB:     "xd",
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoute)
}
