package network

import (
	"chat_controller/config"
	"chat_controller/repository"
	"chat_controller/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg        *config.Config
	engine     *gin.Engine
	service    *service.Service
	repository *repository.Repository
	ip         string
	port       string
}

func NewNetwork(cfg *config.Config, service *service.Service, repository *repository.Repository) *Server {

	s := &Server{
		cfg:        cfg,
		engine:     gin.New(),
		service:    service,
		repository: repository,
		port:       cfg.Server.Port,
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

	return s
}

func (s *Server) Start() error {
	return s.engine.Run(fmt.Sprintf(":%s", s.cfg.Server.Port))
}
