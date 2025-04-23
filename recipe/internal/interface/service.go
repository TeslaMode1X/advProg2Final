package interfaces

import "github.com/TeslaMode1X/advProg2Final/recipe/internal/model"

type RecipeService interface {
	RecipeListService() ([]*model.Recipe, error)
	RecipeCreateService(recipe model.Recipe) (string, error)
	RecipeUpdateService()
	RecipeDeleteService()
}
