package main

import (
	"Tugas_5/delivery/http"
	"Tugas_5/infrastructure"
	"Tugas_5/internal/presenter"
	"Tugas_5/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepo := infrastructure.NewInMemoryUserRepo()
	userUsecase := usecase.NewUserUsecase(userRepo)
	authUsecase := usecase.NewAuthenticationUsecase(userRepo)
	userPresenter := presenter.NewUserPresenter()
	authPresenter := presenter.NewAuthenticationPresenter()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	http.NewUserController(r, userUsecase, authUsecase, userPresenter, authPresenter)

	r.Run(":8080")
}
