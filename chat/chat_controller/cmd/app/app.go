package app

import (
	"chat_controller/config"
	"chat_controller/network"
	"chat_controller/repository"
	"chat_controller/service"
	"log"
)

type App struct {
	cfg        *config.Config
	repository *repository.Repository
	service    *service.Service
	network    *network.Server
}

func NewApp(cfg *config.Config) *App {
	a := &App{cfg: cfg}

	var err error
	if a.repository, err = repository.NewRepository(cfg); err != nil {
		log.Fatalln("failed init repository", "err", err.Error())
	} else {
		a.service = service.NewService(a.repository)
		a.network = network.NewNetwork(cfg, a.service, a.repository)
	}

	return a
}

func (a *App) Start() {
	if err := a.network.Start(); err != nil {
		log.Fatalln("failed start web server", "err", err.Error())
	}
}
