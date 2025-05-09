package repository

import (
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/model"
)

type StatisticsRepository struct {
	db interfaces.Database
}

func NewStatisticsRepository(db interfaces.Database) *StatisticsRepository {
	return &StatisticsRepository{db: db}
}

func (s *StatisticsRepository) GetUsersStatisticsRepo() (*model.UserStatistics, error) {
	const op = "statistics.repository.GetUsersStatistics"

	var userStat *model.UserStatistics

	result := s.db.GetDB().Find(userStat)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return userStat, nil
}

func (s *StatisticsRepository) GetRecipesStatisticsRepo() ([]*model.RecipeReviewStatistics, error) {
	const op = "statistics.repository.GetRecipesStatisticsRepo"

	var recipeStats []*model.RecipeReviewStatistics
	result := s.db.GetDB().Find(&recipeStats)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return recipeStats, nil
}

func (s *StatisticsRepository) GetRecipeStatByIDRepo(id string) (*model.RecipeReviewStatistics, error) {
	const op = "statistics.repository.GetRecipeStatByIDRepo"

	var recipeStat *model.RecipeReviewStatistics

	result := s.db.GetDB().Where("id = ?", id).Find(&recipeStat)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return recipeStat, nil
}
