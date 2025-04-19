package handler

import (
	handler "github.com/TeslaMode1X/advProg2Final/user/internal/handler/response"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model/dto"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) UserRegister(c *gin.Context) {
	const op = "handler.user.UserRegister"

	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Response(c, http.StatusBadRequest, op, err.Error())
		return
	}

	hashedPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		handler.Response(c, http.StatusInternalServerError, op, "Failed to hash password")
		return
	}

	user := model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	id, err := h.userService.UserRegisterService(user)
	if err != nil {
		handler.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	handler.Response(c, http.StatusOK, op, dto.CreateUserResponse{ID: id})
}
