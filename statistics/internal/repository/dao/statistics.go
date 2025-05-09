package dao

import (
	"github.com/gofrs/uuid"
	"time"
)

type UserStatisticsEntity struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	TotalUsers    int       `gorm:"not null"`
	LastUpdatedAt time.Time `gorm:"not null"`
}

func (UserStatisticsEntity) TableName() string {
	return "user_statistics"
}

type RecipeReviewStatisticsEntity struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	RecipeID      uuid.UUID `gorm:"not null;index"`
	TotalReviews  int       `gorm:"not null"`
	TotalRating   float32   `gorm:"not null"`
	AverageRating float32   `gorm:"not null"`
	LastUpdatedAt time.Time `gorm:"not null"`
}

func (RecipeReviewStatisticsEntity) TableName() string {
	return "recipe_reviews"
}
