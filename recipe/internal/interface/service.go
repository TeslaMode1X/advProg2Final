package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/repository/dao"
)

type RecipeService interface {
	RecipeListService() ([]*model.Recipe, error)
	RecipeByIDService(id string) (*dao.RecipeEntity, error)
	RecipeCreateService(recipe model.Recipe) (string, error)
	RecipeUpdateService(recipe model.Recipe) error
	RecipeDeleteService(id, userID string) error
	RecipeExists(id string) (bool, error)
}
