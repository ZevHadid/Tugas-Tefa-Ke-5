package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserPresenterOut struct{}

func NewUserPresenter() *UserPresenterOut {
	return &UserPresenterOut{}
}

func (p *UserPresenterOut) PresentRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
