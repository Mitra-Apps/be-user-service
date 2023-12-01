package service

import (
	"context"
	"user-service/domain/user/entity"
	"user-service/domain/user/repository"
)

type Service struct {
	userRepository repository.UserInterface
}

func New(userRepository repository.UserInterface) *Service {
	return &Service{userRepository: userRepository}
}

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
}