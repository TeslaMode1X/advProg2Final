package dao

import (
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/gofrs/uuid"
	"time"
)

type ReviewEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	RecipeID  uuid.UUID `gorm:"index:idx_recipe"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	Rating    float32   `gorm:"type:float"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (r *ReviewEntity) TableName() string {
	return "review"
}

func ToDao(modelObject *model.Review) *ReviewEntity {
	return &ReviewEntity{
		ID:        modelObject.ID,
		RecipeID:  modelObject.RecipeID,
		UserID:    modelObject.UserID,
		Rating:    modelObject.Rating,
		Comment:   modelObject.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
