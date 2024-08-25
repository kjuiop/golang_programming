package service

import (
	"chat_controller/repository"
	"chat_controller/types/table"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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

	if err := s.repo.Kafka.RegisterSubTopic("chat"); err != nil {
		log.Fatalln("failed topic", "err", err.Error())
	}

	go s.loopSubKafka()

	return s
}

func (s *Service) loopSubKafka() {
	for {
		ev := s.repo.Kafka.Pool(100)

		switch event := ev.(type) {
		case *kafka.Message:

			type ServerInfoEvent struct {
				IP     string
				Status bool
			}

			var decoder ServerInfoEvent
			if err := json.Unmarshal(event.Value, &decoder); err != nil {
				log.Println("failed to decode event", event.Value)
			} else {
				fmt.Println(decoder)
				s.AvgServerList[decoder.IP] = decoder.Status
			}

		case *kafka.Error:
			log.Println("Failed to Pooling event", event.Error())
		}
	}
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
