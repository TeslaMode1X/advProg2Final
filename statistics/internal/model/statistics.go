package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type UserStatistics struct {
	ID            uuid.UUID
	TotalUsers    int
	LastUpdatedAt time.Time
}

// TODO how many reviews some recipe have and average rating
type RecipeReviewStatistics struct {
	ID            uuid.UUID
	RecipeID      uuid.UUID
	TotalReviews  int
	TotalRating   float32
	AverageRating float32
	LastUpdatedAt time.Time
}
