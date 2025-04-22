package main

import (
	"github.com/TeslaMode1X/advProg2Final/api-gateway/internal/handler"
	"github.com/TeslaMode1X/advProg2Final/api-gateway/internal/middleware"
	"github.com/TeslaMode1X/advProg2Final/api-gateway/pkg/load"
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

	err := load.LoadDotEnv()

	userConn, err := grpc.NewClient(userConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to inventory: %v", err)
	}
	defer userConn.Close()

	gatewayHandler := handler.NewGatewayHandler(userConn)

	// USER THING
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", gatewayHandler.UserLogin)
		userGroup.POST("/registration", gatewayHandler.UserRegistration)
		userGroup.GET("/:id", gatewayHandler.UserGetById)

		protected := userGroup.Group("/", middleware.AuthRequired())
		{
			protected.GET("/exists/:id", gatewayHandler.UserExists)
			protected.DELETE("/:id", gatewayHandler.UserDeleteById)
			protected.PUT("", gatewayHandler.UserChangePassword)
		}
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
