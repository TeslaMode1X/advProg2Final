package repository

import (
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
	"github.com/gofrs/uuid"
	"time"
)

type ReviewRepo struct {
	db interfaces.Database
}

func NewReviewRepo(db interfaces.Database) *ReviewRepo {
	return &ReviewRepo{
		db: db,
	}
}

func (r *ReviewRepo) ReviewCreateRepo(model *model.Review) (string, error) {
	const op = "repository.review.ReviewCreateRepo"

	id, _ := uuid.NewV4()

	res := &dao.ReviewEntity{
		ID:        id,
		RecipeID:  model.RecipeID,
		UserID:    model.UserID,
		Rating:    model.Rating,
		Comment:   model.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := r.db.GetDB().Create(&res)
	if err := result.Error; err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return res.ID.String(), nil
}
