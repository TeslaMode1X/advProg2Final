package repository

import (
	"errors"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type UserRepo struct {
	db interfaces.Database
}

func NewUserRepo(db interfaces.Database) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) UserRegisterRepo(user model.User) (uuid.UUID, error) {
	const op = "user.repository.UserRegisterRepo"

	if user.ID == uuid.Nil {
		newID, err := uuid.NewV4()
		if err != nil {
			return uuid.Nil, errors.New(op + ": failed to generate UUID: " + err.Error())
		}
		user.ID = newID
	}

	result := ur.db.GetDB().Create(&user)
	if result.Error != nil {
		return uuid.Nil, errors.New(op + ": failed to create user: " + result.Error.Error())
	}

	return user.ID, nil
}
