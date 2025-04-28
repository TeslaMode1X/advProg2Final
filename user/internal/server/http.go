package server

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/config"
	"github.com/TeslaMode1X/advProg2Final/user/internal/handler"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/user/internal/service"
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

	s.initializeUserHandler()

	serverUrl := fmt.Sprintf(":%s", s.cfg.Server.Port)
	if err := s.app.Run(serverUrl); err != nil {
		s.log.Panic(err)
	}
}

func (s *ginServer) initializeUserHandler() {
	userRepos := repository.NewUserRepo(s.db)
	userService := service.NewUserService(userRepos)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := s.app.Group("/user")
	{
		userRoutes.POST("/registration", userHandler.UserRegister)
		userRoutes.POST("/login", userHandler.UserLogin)
	}

	// TODO apply middleware on some route
	//protected := s.app.Group("/api")
	//protected.Use(middleware.AuthRequired())
	//{
	//	// GET, POST...
	//}
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
