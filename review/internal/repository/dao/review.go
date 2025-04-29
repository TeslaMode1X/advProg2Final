package dao

import (
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
