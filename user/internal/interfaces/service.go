package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type UserService interface {
	UserLoginService(login, password string) (uuid.UUID, error)
	UserRegisterService(user model.User) (uuid.UUID, error)
	//UserExistsService(c *gin.Context)
}
