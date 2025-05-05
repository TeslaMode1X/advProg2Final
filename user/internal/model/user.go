package model

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/repository/dao"
	ownErrors "github.com/TeslaMode1X/advProg2Final/user/pkg/errors"
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

type UserNats struct {
	Email string
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

func ToUserEntityFromUserDomain(u *User) *dao.UserEntity {
	return &dao.UserEntity{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func ToUserDomainFromEntity(ent *dao.UserEntity) *User {
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
