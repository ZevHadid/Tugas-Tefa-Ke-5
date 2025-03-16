package usecase

import (
	"Tugas_5/infrastructure"
	"Tugas_5/internal/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationUsecase struct {
	userRepo infrastructure.UserRepository
}

func NewAuthenticationUsecase(userRepo infrastructure.UserRepository) *AuthenticationUsecase {
	return &AuthenticationUsecase{userRepo: userRepo}
}

func (uc *AuthenticationUsecase) Login(username, email, password string) (string, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil || user.Password != password || user.Email != email {
		return "", err
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &domain.Claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
