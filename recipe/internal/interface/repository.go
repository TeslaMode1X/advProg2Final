package interfaces

import "github.com/TeslaMode1X/advProg2Final/recipe/internal/model"

type RecipeRepository interface {
	RecipeListRepo() ([]*model.Recipe, error)
	RecipeCreateRepo(recipe model.Recipe) (string, error)
	RecipeUpdateRepo()
	RecipeDeleteRepo()
}
