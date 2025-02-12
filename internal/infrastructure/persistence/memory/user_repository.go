package memory

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/0bvim/octoBuddy/config"
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

// fetchFollowers performs an HTTP GET on the provided URL and decodes the JSON array into a slice of User.
func (r *UserRepository) FetchFollowers(token string) ([]entity.Follower, error) {
	var allFollowers []entity.Follower
	client := http.DefaultClient

	for {
		req, err := http.NewRequest("GET", config.FOLLOWERS_URL, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		var followers []entity.Follower
		if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
			return nil, err
		}

		allFollowers = append(allFollowers, followers...)
		if resp.Header.Get("Link") == "" {
			break
		}
	}

	return allFollowers, nil
}