package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	//UserLogin(c *gin.Context)
	UserRegister(c *gin.Context)
	//UserExists(c *gin.Context)
}
