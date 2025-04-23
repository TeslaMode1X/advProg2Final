package server

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/recipe/config"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/handler"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/service"
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
	s.app.Use(gin.Logger())
	s.app.Use(gin.Recovery())

	s.app.GET("/v1/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	s.initializeRecipeHandler()

	serverUrl := fmt.Sprintf(":%s", s.cfg.Server.Port)
	if err := s.app.Run(serverUrl); err != nil {
		s.log.Panic(err)
	}
}

func (s *ginServer) initializeRecipeHandler() {
	recipeRepo := repository.NewRecipeRepo(s.db)
	recipeService := service.NewRecipeService(recipeRepo)
	recipeHandler := handler.NewRecipeHandler(recipeService)

	recipeRoutes := s.app.Group("/recipe")
	{
		recipeRoutes.POST("/create", recipeHandler.RecipeCreate)
		recipeRoutes.GET("/list", recipeHandler.RecipeList)
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
