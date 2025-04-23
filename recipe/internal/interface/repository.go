package interfaces

import "github.com/TeslaMode1X/advProg2Final/recipe/internal/model"

type RecipeRepository interface {
	RecipeListRepo()
	RecipeCreateRepo(recipe model.Recipe) (string, error)
	RecipeUpdateRepo()
	RecipeDeleteRepo()
}
