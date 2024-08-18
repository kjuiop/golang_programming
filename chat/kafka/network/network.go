package network

import (
	"chat-kafka/config"
	"chat-kafka/repository"
	"chat-kafka/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg        *config.Config
	engine     *gin.Engine
	service    *service.Service
	repository *repository.Repository
}

func NewServer(cfg *config.Config, service *service.Service, repository *repository.Repository) *Network {

	n := &Network{
		cfg:        cfg,
		engine:     gin.New(),
		service:    service,
		repository: repository,
	}

	n.engine.Use(gin.Logger())
	n.engine.Use(gin.Recovery())
	n.engine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	r := NewRoom()
	go r.RunInit()

	n.engine.GET("/room", r.SocketServe)

	return n
}

func (n *Network) StartServer() error {
	return n.engine.Run(fmt.Sprintf(":%s", n.cfg.Server.Port))
}
