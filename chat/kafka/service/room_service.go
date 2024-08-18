package service

import "chat-kafka/repository"

type Service struct {
	repo *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repo: repository,
	}
}
