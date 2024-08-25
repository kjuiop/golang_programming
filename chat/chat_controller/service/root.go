package service

import (
	"chat_controller/repository"
	"chat_controller/types/table"
	"log"
)

type Service struct {
	repo *repository.Repository

	AvgServerList map[string]bool
}

func NewService(repository *repository.Repository) *Service {

	s := &Service{
		repo:          repository,
		AvgServerList: make(map[string]bool),
	}

	s.setServerInfo()

	return s
}

func (s *Service) GetAvgServerList() []string {
	var res []string

	for ip, available := range s.AvgServerList {
		if available {
			res = append(res, ip)
		}
	}

	return res
}

func (s *Service) setServerInfo() {

	if serverList, err := s.GetAvailableServerList(); err != nil {
		log.Fatalln("failed set server info", "err", err.Error())
	} else {
		for _, server := range serverList {
			s.AvgServerList[server.IP] = true
		}
	}
}

func (s *Service) GetAvailableServerList() ([]*table.ServerInfo, error) {
	return s.repo.GetAvailableServerList()
}
