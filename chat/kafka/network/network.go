package network

import (
	"chat-kafka/config"
	"chat-kafka/repository"
	"chat-kafka/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg        *config.Config
	engine     *gin.Engine
	service    *service.Service
	repository *repository.Repository
}

func NewServer(cfg *config.Config, service *service.Service, repository *repository.Repository) *Server {

	s := &Server{
		cfg:        cfg,
		engine:     gin.New(),
		service:    service,
		repository: repository,
	}

	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())
	s.engine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	registerServer(s)

	return s
}

func (n *Server) StartServer() error {
	return n.engine.Run(fmt.Sprintf(":%s", n.cfg.Server.Port))
}
