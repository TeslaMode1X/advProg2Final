package server

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/review/config"
	"github.com/TeslaMode1X/advProg2Final/review/internal/handler"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/review/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ginServer struct {
	app *gin.Engine
	db  interfaces.Database
	cfg *config.Config
	log *log.Logger
}

func (s *ginServer) Start() {
	s.app.Use(gin.Recovery())
	s.app.Use(gin.Logger())

	s.app.GET("/v1/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	s.initializeReviewHandler()

	serverUrl := fmt.Sprintf(":%s", s.cfg.Server.Port)
	if err := s.app.Run(serverUrl); err != nil {
		s.log.Panic(err)
	}
}

func (s *ginServer) initializeReviewHandler() {
	reviewRepo := repository.NewReviewRepo(s.db)
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewService)

	reviewGroup := s.app.Group("/review")
	{
		reviewGroup.POST("/create", reviewHandler.ReviewCreateHandler)
		reviewGroup.GET("/list", reviewHandler.ReviewListHandler)
		reviewGroup.GET("/:id", reviewHandler.ReviewByIDHandler)
		reviewGroup.PUT("/update", reviewHandler.ReviewUpdateHandler)
		reviewGroup.DELETE("/delete/:id", reviewHandler.ReviewDeleteHandler)
	}
}

func NewGinServer(conf *config.Config, db interfaces.Database, log *log.Logger) interfaces.Server {
	ginApp := gin.Default()

	return &ginServer{
		app: ginApp,
		db:  db,
		cfg: conf,
		log: log,
	}
}
