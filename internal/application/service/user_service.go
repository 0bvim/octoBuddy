package service

import (
	"github.com/0bvim/octoBuddy/internal/domain/entity"
	"github.com/0bvim/octoBuddy/internal/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUser(id string) (*entity.User, error) {
	return s.userRepo.FindByID(id)
}
