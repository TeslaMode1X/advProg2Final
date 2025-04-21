package main

import (
	"github.com/TeslaMode1X/advProg2Final/api-gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	userConnection = "user:50051"
)

func main() {
	r := gin.Default()

	userConn, err := grpc.NewClient(userConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to inventory: %v", err)
	}
	defer userConn.Close()

	gatewayHandler := handler.NewGatewayHandler(userConn)

	// USER THING
	{
		r.POST("/user/login", gatewayHandler.UserLogin)

		r.POST("/user/registration", gatewayHandler.UserRegistration)

		r.GET("/user/:id", gatewayHandler.UserGetById)

		r.DELETE("/user/:id", gatewayHandler.UserDeleteById)
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
