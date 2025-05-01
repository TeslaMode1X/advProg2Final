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
	userConnection   = "user:50051"
	recipeConnection = "recipe:50052"
	reviewConnection = "review:50053"
)

func main() {
	r := gin.Default()

	err := load.LoadDotEnv()

	userConn, err := grpc.Dial(userConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user: %v", err)
	}
	defer userConn.Close()

	recipeConn, err := grpc.Dial(recipeConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to recipe: %v", err)
	}
	defer recipeConn.Close()

	reviewConn, err := grpc.Dial(reviewConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to review: %v", err)
	}
	defer reviewConn.Close()

	gatewayHandler := handler.NewGatewayHandler(userConn, recipeConn, reviewConn)

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

	// RECIPE THING
	recipeGroup := r.Group("/recipe")
	{
		recipeGroup.GET("/list", gatewayHandler.RecipeList)
		recipeGroup.GET("/:id", gatewayHandler.RecipeByID)
		protected := recipeGroup.Group("/", middleware.AuthRequired())
		{
			protected.POST("/create", gatewayHandler.RecipeCreate)
			protected.PUT("/update", gatewayHandler.RecipeUpdate)
			protected.DELETE("/delete/:id", gatewayHandler.RecipeDelete)
		}
	}

	// REVIEW THING
	reviewGroup := r.Group("/review")
	{
		reviewGroup.GET("/list", gatewayHandler.ReviewList)
		reviewGroup.GET("/:id", gatewayHandler.ReviewById)
		protected := reviewGroup.Group("/", middleware.AuthRequired())
		{
			protected.POST("/create", gatewayHandler.ReviewCreate)
			protected.PUT("/update", gatewayHandler.ReviewUpdate)
			protected.DELETE("/delete/:id", gatewayHandler.ReviewDelete)
		}
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
