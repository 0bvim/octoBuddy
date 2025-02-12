package services

import "github.com/0bvim/octoBuddy/internal/domain/entity"

type AuthService interface {
	GetAuthURL() string
	HandleCallback(code string) (*entity.TokenPair, *entity.User, error)
	RefreshToken(refreshToken string) (*entity.TokenPair, error)
}

type UserService interface {
    GetUser(id string) (*entity.User, error)
	GetUserFollowers(id string) ([]entity.User, error)
	GetUserFollowing(id string) ([]entity.User, error)
}