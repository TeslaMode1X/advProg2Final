package handler

import (
	"github.com/TeslaMode1X/advProg2Final/user/internal/handler/dto"
	handler "github.com/TeslaMode1X/advProg2Final/user/internal/handler/response"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/crypto"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/security"
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

	req.Password = hashedPassword

	user := dto.FromDTO(req)

	id, err := h.userService.UserRegisterService(*user)
	if err != nil {
		handler.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	handler.Response(c, http.StatusOK, op, dto.CreateUserResponse{ID: id})
}

func (h *UserHandler) UserLogin(c *gin.Context) {
	const op = "handler.user.UserLogin"

	var req dto.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Response(c, http.StatusBadRequest, op, err.Error())
		return
	}

	id, err := h.userService.UserLoginService(req.Email, req.Password)
	if err != nil {
		handler.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	token, err := security.CreateToken(id)
	if err != nil {
		handler.Response(c, http.StatusInternalServerError, op, "Failed to generate token")
		return
	}

	c.Header("Authorization", "Bearer "+token)

	handler.Response(c, http.StatusOK, op, dto.CreateUserResponse{ID: id})
}
