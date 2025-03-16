package usecase

import (
	"Tugas-Tefa-Ke-5/infrastructure"
	"Tugas-Tefa-Ke-5/internal/domain"
	"time"
)

type UserUsecase struct {
	userRepo infrastructure.UserRepository
}

func NewUserUsecase(userRepo infrastructure.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uc *UserUsecase) Register(username, email, password string) error {
	user := domain.User{
		ID:        len(uc.userRepo.GetAll()) + 1,
		Name:      username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	return uc.userRepo.Save(user)
}

func (uc *UserUsecase) GetUser(username string) (domain.User, error) {
	return uc.userRepo.FindByUsername(username)
}
