package infrastructure

import (
	"Tugas_5/internal/domain"
	"errors"
	"sync"
)

type UserRepository interface {
	Save(user domain.User) error
	FindByUsername(username string) (domain.User, error)
	GetAll() map[string]domain.User
}

type InMemoryUserRepo struct {
	users map[string]domain.User
	mu    sync.Mutex
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]domain.User),
	}
}

func (r *InMemoryUserRepo) Save(user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.Name] = user
	return nil
}

func (r *InMemoryUserRepo) FindByUsername(username string) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[username]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepo) GetAll() map[string]domain.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.users
}
