package dao

import (
	"github.com/gofrs/uuid"
	"time"
)

type UserEntity struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Username  string     `gorm:"size:50;not null"`
	Password  string     `gorm:"size:255;not null"`
	Email     string     `gorm:"size:254;not null;unique"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

func (u *UserEntity) TableName() string {
	return "users"
}
