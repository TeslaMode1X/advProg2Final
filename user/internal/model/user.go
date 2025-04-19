package model

import (
	ownErrors "github.com/TeslaMode1X/advProg2Final/user/internal/errors"
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) Validate() error {
	if u.Username == "" {
		return ownErrors.ErrorUsernameRequired
	}

	if u.Email == "" {
		return ownErrors.ErrorEmailRequired
	}

	return nil
}

func ToUserEntityFromUserDomain(u *User) *UserEntity {
	return &UserEntity{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

type UserEntity struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username  string     `gorm:"size:50;not null"`
	Password  string     `gorm:"size:255;not null"`
	Email     string     `gorm:"size:254;not null;unique"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

func (u *UserEntity) TableName() string {
	return "user"
}

func ToUserDomainFromEntity(ent *UserEntity) *User {
	return &User{
		ID:        ent.ID,
		Username:  ent.Username,
		Password:  ent.Password,
		Email:     ent.Email,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
		DeletedAt: ent.DeletedAt,
	}
}
