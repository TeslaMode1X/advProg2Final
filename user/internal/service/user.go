package service

import (
	"errors"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type UserService struct {
	userRepo interfaces.UserRepo
}

func NewUserService(userRepo interfaces.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) UserRegisterService(user model.User) (uuid.UUID, error) {
	const op = "user.service.UserRegisterService"

	if err := user.Validate(); err != nil {
		return uuid.Nil, err
	}

	id, err := us.userRepo.UserRegisterRepo(user)
	if err != nil {
		return uuid.Nil, errors.New(op + ": " + err.Error())
	}

	return id, nil
}

func (us *UserService) UserLoginService(login, password string) (uuid.UUID, error) {
	const op = "user.service.UserLoginService"

	id, err := us.userRepo.UserLoginRepo(login, password)
	if err != nil {
		return uuid.Nil, errors.New(op + ": " + err.Error())
	}

	return id, nil
}

func (us *UserService) UserGetByIdService(id string) (*model.User, error) {
	const op = "service.user.UserGetByIdService"

	user, err := us.userRepo.UserGetByIdRepo(id)
	if err != nil {
		return nil, errors.New(op + ": " + err.Error())
	}

	return user, nil
}

func (us *UserService) UserDeleteByIdService(id string) error {
	const op = "service.user.UserDeleteByIdService"

	if err := us.userRepo.UserDeleteByIdRepo(id); err != nil {
		return errors.New(op + ": " + err.Error())
	}

	return nil
}
