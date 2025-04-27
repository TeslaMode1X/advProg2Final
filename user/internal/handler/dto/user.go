package dto

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

func ToUserDomainFromUserRequest(req CreateUserRequest) (*model.User, error) {
	user := &model.User{
		ID:        uuid.Must(uuid.NewV4()),
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}
