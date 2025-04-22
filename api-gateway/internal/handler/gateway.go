package handler

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/user"
	"github.com/gofrs/uuid"

	"google.golang.org/grpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	userClient user.UserServiceClient
}

func NewGatewayHandler(userConn *grpc.ClientConn) *GatewayHandler {
	return &GatewayHandler{
		userClient: user.NewUserServiceClient(userConn),
	}
}

func (g *GatewayHandler) UserLogin(c *gin.Context) {
	_, err := c.Cookie("auth_token")
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "already logged in"})
		return
	}

	var req user.RequestUserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	resp, err := g.userClient.UserLogin(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		"auth_token",
		resp.Token,
		3600*1,
		"/",
		"",
		false,
		false,
	)

	c.JSON(http.StatusOK, gin.H{
		"id":    resp.Id,
		"token": resp.Token,
	})
}

func (g *GatewayHandler) UserRegistration(c *gin.Context) {
	var req user.RequestUserRegistration

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Email == "" || req.Password == "" || req.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, email and password are required"})
		return
	}

	resp, err := g.userClient.UserRegistration(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": resp.Id})
}

func (g *GatewayHandler) UserGetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &user.RequestUserGetById{
		Id: id,
	}

	user, err := g.userClient.UserGetById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (g *GatewayHandler) UserDeleteById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &user.RequestUserGetById{
		Id: id,
	}

	_, err := g.userClient.UserDeleteById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": "was deleted"})
}

func (g *GatewayHandler) UserExists(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &user.RequestUserGetById{
		Id: id,
	}

	exists, err := g.userClient.UserExists(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": exists})
}

func (g *GatewayHandler) UserChangePassword(c *gin.Context) {
	var req user.RequestUserChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Old and new passwords are required"})
		return
	}

	userIDStr := userID.(uuid.UUID).String()

	ctx := context.Background()
	_, err := g.userClient.UserChangePassword(ctx, &user.RequestUserChangePassword{
		Id:          userIDStr,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
