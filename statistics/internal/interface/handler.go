package interfaces

import "github.com/gin-gonic/gin"

type StatisticsHandler interface {
	GetUsersStatistics(c *gin.Context)
	GetRecipesStatistics(c *gin.Context)
	GetRecipeStatByID(c *gin.Context)
}
