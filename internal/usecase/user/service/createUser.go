package service

import (
	"context"
	"lockStock/internal/domain/user"
	"lockStock/internal/domain/user/contract"
)

type UserService struct {
	repo contract.UserRepository
}

func NewUserService(repo contract.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, newUser *user.User) (string, error) {
	return s.repo.SaveUser(ctx, newUser)
}
