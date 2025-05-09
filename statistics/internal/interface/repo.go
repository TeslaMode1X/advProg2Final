package interfaces

import "github.com/TeslaMode1X/advProg2Final/statistics/internal/model"

type StatisticsRepo interface {
	GetUsersStatisticsRepo() (*model.UserStatistics, error)
	GetRecipesStatisticsRepo() ([]*model.RecipeReviewStatistics, error)
	GetRecipeStatByIDRepo(id string) (*model.RecipeReviewStatistics, error)
	AddNewUserCounter() error
}
