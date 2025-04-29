package dto

import (
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/gofrs/uuid"
)

type ReviewCreateRequest struct {
	RecipeID string  `json:"recipeId"`
	UserID   string  `json:"userId"`
	Rating   float32 `json:"rating"`
	Comment  string  `json:"comment"`
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
