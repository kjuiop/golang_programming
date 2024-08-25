package service

import (
	"chat-kafka/repository"
	"chat-kafka/types/schema"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Service struct {
	repo *repository.Repository
}

type ServerInfoEvent struct {
	IP     string
	Status bool
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) ServerSet(ip string, available bool) error {
	if err := s.repo.ServerSet(ip, available); err != nil {
		log.Println("Failed To ServerSet", "ip", ip, "available", available)
		return err
	} else {
		return nil
	}
}

func (s *Service) PublishServerStatusEvent(addr string, status bool) {

	// kafka 이벤트 전송
	e := &ServerInfoEvent{IP: addr, Status: status}
	ch := make(chan kafka.Event)

	if v, err := json.Marshal(e); err != nil {
		log.Println("Failed to Marshal")
	} else if result, err := s.PublishEvent("chat", v, ch); err != nil {
		log.Println("Failed To Send Event To Kafka", "err", err)
	} else {
		// Send Event To Kafka
		log.Println("Success To Send Event", result)
	}

}

func (s *Service) PublishEvent(topic string, value []byte, ch chan kafka.Event) (kafka.Event, error) {
	return s.repo.Kafka.PublishEvent(topic, value, ch)
}

func (s *Service) EnterRoom(roomName string) ([]*schema.Chat, error) {
	if res, err := s.repo.GetChatList(roomName); err != nil {
		log.Println("Failed To Get All Chat List", "err", err.Error())
		return nil, err
	} else {
		return res, nil
	}
}

func (s *Service) RoomList() ([]*schema.Room, error) {
	if res, err := s.repo.RoomList(); err != nil {
		log.Println("Failed To Get All Room List", "err", err.Error())
		return nil, err
	} else {
		return res, nil
	}
}

func (s *Service) MakeRoom(name string) error {
	if err := s.repo.MakeRoom(name); err != nil {
		log.Println("Failed To Make Room", "err", err.Error())
		return err
	} else {
		return nil
	}
}

func (s *Service) Room(name string) (*schema.Room, error) {
	if res, err := s.repo.Room(name); err != nil {
		log.Println("Failed To Get Room", "err", err.Error())
		return nil, err
	} else {
		return res, nil
	}
}

func (s *Service) InsertChatting(user string, message string, roomName string) {
	if err := s.repo.InsertChatting(user, message, roomName); err != nil {
		log.Printf("failed insert chatting, err : %v\n", err)
	}
}
