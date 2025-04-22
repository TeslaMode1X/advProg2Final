package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID
	RecipeID  uuid.UUID
	UserID    uuid.UUID
	Rating    int
	Comment   string
	CreatedAt time.Time
}
