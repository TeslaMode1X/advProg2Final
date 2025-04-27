package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type Recipe struct {
	ID            uuid.UUID
	AuthorID      uuid.UUID
	Title         string
	Description   string
	Photos        []string
	AverageRating float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
