package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type UserService interface {
	UserLoginService(login, password string) (uuid.UUID, error)
	UserRegisterService(user model.User) (uuid.UUID, error)
	UserGetByIdService(id string) (*model.User, error)
	UserDeleteByIdService(id string) error
	UserExistsService(id string) (bool, error)
	UserUpdatePasswordService(id, oldPassword, newPassword string) error
}
