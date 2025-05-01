package repository

import (
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
	"github.com/gofrs/uuid"
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

	model.ID = id

	res := dao.ToDao(model)

	result := r.db.GetDB().Create(&res)
	if err := result.Error; err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return res.ID.String(), nil
}

func (r *ReviewRepo) ReviewListRepo() ([]*dao.ReviewEntity, error) {
	const op = "repository.review.ReviewListRepo"

	var modelObjects []*dao.ReviewEntity

	result := r.db.GetDB().Find(&modelObjects)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return modelObjects, nil
}
