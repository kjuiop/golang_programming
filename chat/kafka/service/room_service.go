package service

import (
	"chat-kafka/repository"
	"chat-kafka/types/schema"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Service struct {
	repo *repository.Repository
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