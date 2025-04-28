package server

import (
	"github.com/TeslaMode1X/advProg2Final/proto/gen/recipe"
	"github.com/TeslaMode1X/advProg2Final/recipe/config"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/service"
	grpcService "github.com/TeslaMode1X/advProg2Final/recipe/internal/service/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcServerObject struct {
	server *grpc.Server
	cfg    *config.Config
	db     interfaces.Database
	log    *log.Logger
}

func NewGrpcServer(conf *config.Config, db interfaces.Database, log *log.Logger) interfaces.Server {
	recipeRepository := repository.NewRecipeRepo(db)
	recipeService := service.NewRecipeService(recipeRepository)

	grpcServer := grpc.NewServer()

	recipe.RegisterRecipeServiceServer(grpcServer, grpcService.NewRecipeServerGrpc(recipeService))

	return &grpcServerObject{
		server: grpcServer,
		cfg:    conf,
		db:     db,
		log:    log,
	}
}

func (s *grpcServerObject) Start() {
	port := ":50052"
	if s.cfg.Server.Port != "" {
		port = ":" + s.cfg.Server.Port
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		s.log.Fatalf("Failed to listen on %s: %v", port, err)
	}

	s.log.Printf("Starting recipe gRPC server on %s", port)
	if err = s.server.Serve(lis); err != nil {
		s.log.Fatalf("Failed to serve: %v", err)
	}
}
