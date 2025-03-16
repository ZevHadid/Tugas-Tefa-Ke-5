package presenter

import "github.com/gin-gonic/gin"

type UserPresenterIn interface {
	PresentRegister(c *gin.Context)
}
