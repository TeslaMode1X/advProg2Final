package interfaces

import "github.com/gin-gonic/gin"

type ReviewHandler interface {
	ReviewCreateHandler(c *gin.Context)
	ReviewListHandler(c *gin.Context)
	ReviewByIDHandler(c *gin.Context)
	ReviewUpdateHandler(c *gin.Context)
	ReviewDeleteHandler(c *gin.Context)
}
