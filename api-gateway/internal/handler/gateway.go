package handler

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/user"
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

	c.Header("Authorization", "Bearer "+resp.Token)
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
