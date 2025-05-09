package server

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/statistics/config"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/handler"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/service"
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

	s.initializeStatisticsHandler()

	serverUrl := fmt.Sprintf(":%s", s.cfg.Server.Port)
	if err := s.app.Run(serverUrl); err != nil {
		s.log.Panic(err)
	}
}

func (s *ginServer) initializeStatisticsHandler() {
	repositoryStats := repository.NewStatisticsRepository(s.db)
	serviceStats := service.NewStatisticsService(repositoryStats)
	serviceHandler := handler.NewStatisticsHandler(serviceStats)

	statsGroup := s.app.Group("/statistics")
	{
		statsGroup.GET("/users", serviceHandler.GetUsersStatistics)
		statsGroup.GET("/recipe", serviceHandler.GetRecipesStatistics)
		statsGroup.GET("/recipe/:id", serviceHandler.GetRecipeStatByID)
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
