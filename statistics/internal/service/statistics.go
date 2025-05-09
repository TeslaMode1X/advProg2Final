package service

import (
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/model"
)

type StatisticsService struct {
	repo interfaces.StatisticsRepo
}

func NewStatisticsService(repo interfaces.StatisticsRepo) *StatisticsService {
	return &StatisticsService{repo: repo}
}

func (s *StatisticsService) GetUsersStatisticsService() (*model.UserStatistics, error) {
	const op = "statistics.service.GetUsersStatisticsService"

	userStats, err := s.repo.GetUsersStatisticsRepo()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return userStats, nil
}

func (s *StatisticsService) GetRecipesStatisticsService() ([]*model.RecipeReviewStatistics, error) {
	const op = "statistics.service.GetRecipesStatisticsService"

	recipeStats, err := s.repo.GetRecipesStatisticsRepo()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return recipeStats, nil
}

func (s *StatisticsService) GetRecipeStatByIDService(id string) (*model.RecipeReviewStatistics, error) {
	const op = "statistics.service.GetRecipeStatByIDService"

	recipeStat, err := s.repo.GetRecipeStatByIDRepo(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return recipeStat, nil
}

func (s *StatisticsService) AddNewUserCounter() error {
	const op = "statistics.service.AddNewUserCounter"

	err := s.repo.AddNewUserCounter()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
