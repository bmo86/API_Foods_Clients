package main

import (
	"foods_API_GRPC/database"
	"foods_API_GRPC/proto/foodspb"
	"foods_API_GRPC/serverpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewConnectionDatabase("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	serverpb := serverpb.NewServerpb(repo)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	foodspb.RegisterFoodServiceServer(s, serverpb)
	reflection.Register(s)
	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}

}
