package interfaces

import "github.com/gin-gonic/gin"

type RecipeHandler interface {
	RecipeList(c *gin.Context)
	RecipeByID(c *gin.Context)
	RecipeCreate(c *gin.Context)
	RecipeUpdate(c *gin.Context)
	RecipeDelete(c *gin.Context)
}
