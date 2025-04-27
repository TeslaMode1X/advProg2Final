package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/repository/dao"
)

type RecipeRepository interface {
	RecipeListRepo() ([]*model.Recipe, error)
	RecipeByIDRepo(id string) (*dao.RecipeEntity, error)
	RecipeCreateRepo(recipe model.Recipe) (string, error)
	RecipeUpdateRepo(recipe model.Recipe) error
	RecipeDeleteRepo(id string) error
}
