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

func (r *ReviewRepo) ReviewByIDRepo(id string) (*dao.ReviewEntity, error) {
	const op = "repository.review.ReviewByIdRepo"

	var object dao.ReviewEntity

	result := r.db.GetDB().Where("id = ?", id).First(&object)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &object, nil
}

func (r *ReviewRepo) ReviewUpdateRepo(model *model.Review) error {
	const op = "repository.review.ReviewUpdateRepo"

	var dao dao.ReviewEntity

	findResult := r.db.GetDB().Where("id = ?", model.ID).First(&dao)
	if err := findResult.Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	dao.RecipeID = model.RecipeID
	dao.UserID = model.UserID
	dao.Rating = model.Rating
	dao.Comment = model.Comment

	saveResult := r.db.GetDB().Save(&dao)
	if err := saveResult.Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ReviewRepo) ReviewDeleteRepo(id string) error {
	const op = "repository.review.ReviewDeleteRepo"

	var dao dao.ReviewEntity
	result := r.db.GetDB().Where("id = ?", id).Delete(&dao)
	if err := result.Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ReviewRepo) ReviewUserCheck(userId, reviewId string) (bool, error) {
	const op = "repository.review.ReviewUserCheck"

	var dao dao.ReviewEntity
	result := r.db.GetDB().Where("id = ? and user_id = ?", reviewId, userId).First(&dao)
	if err := result.Error; err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, nil
}
