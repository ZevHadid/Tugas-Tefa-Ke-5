package http

import (
	"net/http"

	"Tugas_5/internal/domain"
	"Tugas_5/internal/presenter"
	"Tugas_5/internal/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase   *usecase.UserUsecase
	authUsecase   *usecase.AuthenticationUsecase
	userPresenter *presenter.UserPresenterOut
	authPresenter *presenter.AuthenticationPresenterOut
}

func NewUserController(r *gin.Engine, userUsecase *usecase.UserUsecase, authUsecase *usecase.AuthenticationUsecase, userPresenter *presenter.UserPresenterOut, authPresenter *presenter.AuthenticationPresenterOut) {
	controller := &UserController{
		userUsecase:   userUsecase,
		authUsecase:   authUsecase,
		userPresenter: userPresenter,
		authPresenter: authPresenter,
	}

	r.GET("/register", controller.userPresenter.PresentRegister)
	r.POST("/register", controller.Register)
	r.GET("/login", controller.authPresenter.PresentLogin)
	r.POST("/login", controller.Login)
	r.GET("/profile", controller.Profile)
	r.POST("/logout", controller.Logout)
}

func (c *UserController) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if username == "" || password == "" || email == "" {
		ctx.String(http.StatusBadRequest, "Invalid input")
		return
	}
	err := c.userUsecase.Register(username, email, password)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error registering user")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/login")
}

func (c *UserController) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	tokenString, err := c.authUsecase.Login(username, email, password)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Invalid credentials")
		return
	}
	ctx.SetCookie("token", tokenString, 300, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/profile")
}

func (c *UserController) Profile(ctx *gin.Context) {
	tokenStr, err := ctx.Cookie("token")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	claims := &domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	user, err := c.userUsecase.GetUser(claims.Username)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	ctx.HTML(http.StatusOK, "profile.html", gin.H{
		"username":  user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt.Format("2006-01-02"),
	})
}

func (c *UserController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/login")
}
