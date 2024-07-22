package service

import (
	"context"
	"learning/grpc/pkg/model"

	"github.com/samber/lo"
)

type userRepository interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

type userService struct {
}

func NewUserService() *userService {
	return &userService{
	}
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return &model.User{
		ID: 1,
		DisplayName: lo.ToPtr("name"),
		Username: lo.ToPtr("username"),
		Email: lo.ToPtr("abc@gmail.com"),
	}, nil
}
