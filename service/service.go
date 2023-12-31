package service

import (
	"context"

	"github.com/Mitra-Apps/be-user-service/domain/user/entity"
	"github.com/Mitra-Apps/be-user-service/domain/user/repository"
)

type Service struct {
	userRepository repository.UserInterface
}

func New(userRepository repository.UserInterface) *Service {
	return &Service{userRepository: userRepository}
}

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	Login(ctx context.Context, payload entity.LoginRequest) (*entity.User, error)
}
