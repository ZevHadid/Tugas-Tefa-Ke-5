package presenter

import "github.com/gin-gonic/gin"

type AuthenticationPresenterIn interface {
	PresentLogin(c *gin.Context)
}
