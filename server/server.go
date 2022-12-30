package server

import (
	"context"
	"errors"
	"foods_API_GRPC/websocket"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type Config struct {
	Port      string
	UrlDB     string
	JWTSecret string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *gin.Engine
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}

	if config.UrlDB == "" {
		return nil, errors.New("database is required")
	}

	b := &Broker{
		config: config,
		router: gin.New(),
		hub:    websocket.NewHub(),
	}

	return b, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
	b.router = gin.New()
	binder(b, b.router)

	handler := cors.Default().Handler(b.router)

	//datebase

	go b.Hub().Run()

	//set repo database

	log.Println("server on port : ", b.config.Port)
	if err := http.ListenAndServe(":"+b.config.Port, handler); err != nil {
		log.Fatal("List And Server - error type : ", err)
	}

}
