package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationPresenterOut struct{}

func NewAuthenticationPresenter() *AuthenticationPresenterOut {
	return &AuthenticationPresenterOut{}
}

func (p *AuthenticationPresenterOut) PresentLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
