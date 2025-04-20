package repository

import (
	"errors"
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/crypto"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
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

func (ur *UserRepo) UserLoginRepo(login, password string) (uuid.UUID, error) {
	const op = "user.repository.UserLoginRepo"

	var user model.User
	result := ur.db.GetDB().Where("email = ?", login).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return uuid.Nil, fmt.Errorf("%s: user witch such email %s not found", op, login)
		}
		return uuid.Nil, fmt.Errorf("%s: %w", op, result.Error)
	}

	if !crypto.VerifyPassword(user.Password, password) {
		return uuid.Nil, fmt.Errorf("%s: wrong password", op)
	}

	return user.ID, nil
}
