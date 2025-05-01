package dto

import (
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
	errors "github.com/TeslaMode1X/advProg2Final/review/pkg"
	"github.com/gofrs/uuid"
)

type ReviewCreateRequest struct {
	RecipeID string  `json:"recipeId"`
	UserID   string  `json:"userId"`
	Rating   float32 `json:"rating"`
	Comment  string  `json:"comment"`
}

func (r *ReviewCreateRequest) Validate() error {
	if r.RecipeID == "" {
		return errors.ErrRecipeIdEmpty
	}

	if r.UserID == "" {
		return errors.ErrUserIdEmpty
	}

	if r.Rating < 0.0 || r.Rating > 5.0 {
		return errors.ErrInvalidRating
	}

	if r.Comment == "" {
		return errors.ErrCommentEmpty
	}

	return nil
}

func FromDTO(modelObject ReviewCreateRequest) *model.Review {
	recipeUUID, _ := uuid.FromString(modelObject.RecipeID)
	userUUID, _ := uuid.FromString(modelObject.UserID)

	return &model.Review{
		RecipeID: recipeUUID,
		UserID:   userUUID,
		Rating:   modelObject.Rating,
		Comment:  modelObject.Comment,
	}
}

func ToDTO(reviews []*dao.ReviewEntity) []*ReviewCreateRequest {
	var list []*ReviewCreateRequest
	for _, review := range reviews {
		var object ReviewCreateRequest
		object.RecipeID = review.RecipeID.String()
		object.UserID = review.UserID.String()
		object.Rating = review.Rating
		object.Comment = review.Comment

		list = append(list, &object)
	}

	return list
}
