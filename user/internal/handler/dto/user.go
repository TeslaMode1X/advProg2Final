package dto

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

func FromDTO(modelObject CreateUserRequest) *model.User {
	return &model.User{
		Username: modelObject.Username,
		Password: modelObject.Password,
		Email:    modelObject.Email,
	}
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}
