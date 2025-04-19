package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type UserRepo interface {
	//UserLoginRepo()
	UserRegisterRepo(user model.User) (uuid.UUID, error)
}
