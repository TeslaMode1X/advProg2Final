package repository

import (
	"context"
	"errors"
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/model"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type StatisticsRepository struct {
	db interfaces.Database
}

func NewStatisticsRepository(db interfaces.Database) *StatisticsRepository {
	return &StatisticsRepository{db: db}
}

func (s *StatisticsRepository) GetUsersStatisticsRepo() (*model.UserStatistics, error) {
	const op = "statistics.repository.GetUsersStatistics"

	userStat := &model.UserStatistics{}

	result := s.db.GetDB().First(userStat)

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

func (s *StatisticsRepository) AddNewUserCounter() error {
	const op = "statistics.repository.AddNewUserCounter"

	var userStat model.UserStatistics

	result := s.db.GetDB().First(&userStat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newID, err := uuid.NewV4()
			if err != nil {
				return fmt.Errorf("%s: failed to generate UUID: %w", op, err)
			}

			userStat = model.UserStatistics{
				ID:            newID,
				TotalUsers:    1,
				LastUpdatedAt: time.Now(),
			}

			if err := s.db.GetDB().Create(&userStat).Error; err != nil {
				return fmt.Errorf("%s: failed to create new user statistics: %w", op, err)
			}

			s.db.GetDB().Logger.Info(context.Background(), "Created new user statistics record with ID: %s", newID.String())
			return nil
		}

		return fmt.Errorf("%s: failed to query user statistics: %w", op, result.Error)
	}

	userStat.TotalUsers++
	userStat.LastUpdatedAt = time.Now()

	if err := s.db.GetDB().Save(&userStat).Error; err != nil {
		return fmt.Errorf("%s: failed to update user statistics: %w", op, err)
	}

	return nil
}
