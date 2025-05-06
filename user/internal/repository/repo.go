package repository

import (
	"errors"
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/crypto"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"log"
)

type UserRepo struct {
	db interfaces.Database
}

func NewUserRepo(db interfaces.Database) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) UserRegisterRepo(user model.User) (uuid.UUID, error) {
	const op = "user.repository.UserRegisterRepo"

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

func (ur *UserRepo) UserGetByIdRepo(id string) (*model.User, error) {
	const op = "user.repository.UserGetByIdRepo"

	var user model.User
	result := ur.db.GetDB().Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: user with such id %s not found", op, id)
		}
		return nil, fmt.Errorf("%s: %w", op, result.Error)
	}

	return &user, nil
}

func (ur *UserRepo) UserDeleteByIdRepo(id string) error {
	const op = "user.repository.UserDeleteByIdRepo"

	result := ur.db.GetDB().Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%s: user with such id %s not found", op, id)
		}
		return fmt.Errorf("%s: %w", op, result.Error)
	}

	return nil
}

func (ur *UserRepo) UserExistsRepo(id string) (bool, error) {
	const op = "user.repository.UserExistsRepo"

	var user model.User
	result := ur.db.GetDB().Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("%s: user with such id %s not found", op, id)
		}
		return false, fmt.Errorf("%s: %w", op, result.Error)
	}

	return true, nil
}

func (ur *UserRepo) UserUpdatePasswordRepo(id, oldPassword, newPassword string) error {
	const op = "grpc.repository.UserUpdatePasswordRepo"

	var user model.User
	if err := ur.db.GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("%s: user not found: %v", op, err)
			return errors.New("user not found")
		}
		log.Printf("%s: failed to query user: %v", op, err)
		return errors.New("failed to retrieve user")
	}

	if err := crypto.VerifyPassword(user.Password, oldPassword); err == false {
		log.Printf("%s: incorrect old password: %v", op, err)
		return errors.New("incorrect old password")
	}

	hashedNewPassword, err := crypto.HashPassword(newPassword)
	if err != nil {
		log.Printf("%s: failed to hash new password: %v", op, err)
		return errors.New("failed to hash new password")
	}

	user.Password = hashedNewPassword
	if err := ur.db.GetDB().Save(&user).Error; err != nil {
		log.Printf("%s: failed to update password: %v", op, err)
		return errors.New("failed to update password")
	}

	return nil
}
