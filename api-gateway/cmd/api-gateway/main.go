package main

import (
	"github.com/TeslaMode1X/advProg2Final/api-gateway/internal/handler"
	"github.com/TeslaMode1X/advProg2Final/api-gateway/internal/middleware"
	"github.com/TeslaMode1X/advProg2Final/api-gateway/pkg/load"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	userConnection       = "user:50051"
	recipeConnection     = "recipe:50052"
	reviewConnection     = "review:50053"
	statisticsConnection = "statistics:50054"

	//Metrics
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_request_total",
			Help: "Total Number Of requests to API Gateway",
		},
		[]string{"path"})
)

func init() {
	//METRICS REGISTRATION!
	prometheus.MustRegister(requestCounter)
}

func main() {
	r := gin.Default()

	//Prometheus endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

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

	statisticsConn, err := grpc.Dial(statisticsConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to review: %v", err)
	}
	defer statisticsConn.Close()

	gatewayHandler := handler.NewGatewayHandler(userConn, recipeConn, reviewConn, statisticsConn)

	// USER THING
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", func(c *gin.Context) {
			requestCounter.WithLabelValues("/user/login").Inc()
			gatewayHandler.UserLogin(c)
		})
		userGroup.POST("/registration", func(c *gin.Context) {
			requestCounter.WithLabelValues("/user/registration").Inc()
			gatewayHandler.UserRegistration(c)
		})
		userGroup.GET("/:id", func(c *gin.Context) {
			requestCounter.WithLabelValues("/user/:id").Inc()
			gatewayHandler.UserGetById(c)
		})

		protected := userGroup.Group("/", middleware.AuthRequired())
		{
			protected.GET("/exists/:id", func(c *gin.Context) {
				requestCounter.WithLabelValues("/user/exists/:id").Inc()
				gatewayHandler.UserExists(c)
			})
			protected.DELETE("/:id", func(c *gin.Context) {
				requestCounter.WithLabelValues("/user/:id").Inc()
				gatewayHandler.UserDeleteById(c)
			})
			protected.PUT("", func(c *gin.Context) {
				requestCounter.WithLabelValues("/user").Inc()
				gatewayHandler.UserChangePassword(c)
			})
		}
	}

	// RECIPE THING
	recipeGroup := r.Group("/recipe")
	{
		recipeGroup.GET("/list", func(c *gin.Context) {
			requestCounter.WithLabelValues("/recipe/list").Inc()
			gatewayHandler.RecipeList(c)
		})
		recipeGroup.GET("/:id", func(c *gin.Context) {
			requestCounter.WithLabelValues("/recipe/:id").Inc()
			gatewayHandler.RecipeByID(c)
		})
		protected := recipeGroup.Group("/", middleware.AuthRequired())
		{
			protected.POST("/create", func(c *gin.Context) {
				requestCounter.WithLabelValues("/recipe/create").Inc()
				gatewayHandler.RecipeCreate(c)
			})
			protected.PUT("/update", func(c *gin.Context) {
				requestCounter.WithLabelValues("/recipe/update").Inc()
				gatewayHandler.RecipeUpdate(c)
			})
			protected.DELETE("/delete/:id", func(c *gin.Context) {
				requestCounter.WithLabelValues("/recipe/delete/:id").Inc()
				gatewayHandler.RecipeDelete(c)
			})
		}
	}

	// REVIEW THING
	reviewGroup := r.Group("/review")
	{
		reviewGroup.GET("/list", func(c *gin.Context) {
			requestCounter.WithLabelValues("/review/list").Inc()
			gatewayHandler.ReviewList(c)
		})
		reviewGroup.GET("/:id", func(c *gin.Context) {
			requestCounter.WithLabelValues("/review/:id").Inc()
			gatewayHandler.ReviewById(c)
		})
		protected := reviewGroup.Group("/", middleware.AuthRequired())
		{
			protected.POST("/create", func(c *gin.Context) {
				requestCounter.WithLabelValues("/review/create").Inc()
				gatewayHandler.ReviewCreate(c)
			})
			protected.PUT("/update", func(c *gin.Context) {
				requestCounter.WithLabelValues("/review/update").Inc()
				gatewayHandler.ReviewUpdate(c)
			})
			protected.DELETE("/delete/:id", func(c *gin.Context) {
				requestCounter.WithLabelValues("/review/delete/:id").Inc()
				gatewayHandler.ReviewDelete(c)
			})
		}
	}

	// STATISTICS THING
	statisticsGroup := r.Group("/statistics")
	{
		statisticsGroup.GET("/users", func(c *gin.Context) {
			requestCounter.WithLabelValues("/statistics/users").Inc()
			gatewayHandler.GetUserRegisteredStatistics(c)
		})
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
