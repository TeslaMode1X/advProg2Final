package dao

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"time"
)

type RecipeEntity struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	AuthorID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Title         string    `gorm:"size:100;not null"`
	Description   string    `gorm:"type:text"`
	Photos        datatypes.JSON
	AverageRating float32   `gorm:"type:decimal(3,2);default:0.0"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (r *RecipeEntity) TableName() string {
	return "recipe"
}
