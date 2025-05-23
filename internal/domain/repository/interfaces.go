package repository

import "github.com/0bvim/octoBuddy/internal/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	FetchFollowers(token string) ([]entity.Follower, error)
}
