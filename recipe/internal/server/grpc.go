package server

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/recipe"
	"github.com/TeslaMode1X/advProg2Final/recipe/config"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/service"
	grpcService "github.com/TeslaMode1X/advProg2Final/recipe/internal/service/grpc"
	"github.com/TeslaMode1X/advProg2Final/recipe/pkg/nats"
	"github.com/TeslaMode1X/advProg2Final/recipe/pkg/nats/producer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcServerObject struct {
	server     *grpc.Server
	cfg        *config.Config
	db         interfaces.Database
	log        *log.Logger
	natsClient *nats.Client
}

func NewGrpcServer(conf *config.Config, db interfaces.Database, log *log.Logger) interfaces.Server {
	recipeRepository := repository.NewRecipeRepo(db)
	recipeService := service.NewRecipeService(recipeRepository)

	// NATS connection
	natsClient, err := nats.NewClient(context.Background(), []string{"nats_service:4222"}, "", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	// Sending objects via NATS
	recipeProducer := producer.NewRecipeProducer(natsClient)

	grpcServer := grpc.NewServer()

	recipe.RegisterRecipeServiceServer(grpcServer, grpcService.NewRecipeServerGrpc(recipeService, recipeProducer))

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
