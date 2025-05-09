package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/model"
)

type StatisticsService interface {
	GetUsersStatisticsService() (*model.UserStatistics, error)
	GetRecipesStatisticsService() ([]*model.RecipeReviewStatistics, error)
	GetRecipeStatByIDService(id string) (*model.RecipeReviewStatistics, error)
	AddNewUserCounter() error
}
