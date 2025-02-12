package memory

import (
	"strconv"
	"sync"

	"github.com/0bvim/octoBuddy/internal/domain/entity"
)

type UserRepository struct {
	users map[string]*entity.User
	mu    sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *UserRepository) Save(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[strconv.Itoa(user.ID)] = user
	return nil
}

func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}
