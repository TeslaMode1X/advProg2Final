package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type UserRepo interface {
	UserLoginRepo(login, password string) (uuid.UUID, error)
	UserRegisterRepo(user model.User) (uuid.UUID, error)
	UserGetByIdRepo(id string) (*model.User, error)
	UserGetAllRepo() ([]*model.User, error)
	UserDeleteByIdRepo(id string) error
	UserExistsRepo(id string) (bool, error)
	UserUpdatePasswordRepo(id, oldPassword, newPassword string) error
}
