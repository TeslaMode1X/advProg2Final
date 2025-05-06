package server

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/review"
	"github.com/TeslaMode1X/advProg2Final/review/config"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/review/internal/service"
	grpcService "github.com/TeslaMode1X/advProg2Final/review/internal/service/grpc"
	"github.com/TeslaMode1X/advProg2Final/review/pkg/nats"
	"github.com/TeslaMode1X/advProg2Final/review/pkg/nats/producer"

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

func NewGrpcServer(cfg *config.Config, db interfaces.Database, log *log.Logger) *grpcServerObject {
	reviewRepository := repository.NewReviewRepo(db)
	reviewService := service.NewReviewService(reviewRepository)

	ctx := context.Background()

	// NATS connection
	natsClient, err := nats.NewClient(ctx, []string{"nats_service:4222"}, "", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	// Sending objects via NATS
	reviewProducer := producer.NewReviewProducer(natsClient)

	grpcServer := grpc.NewServer()

	review.RegisterReviewServiceServer(grpcServer, grpcService.NewReviewServerGrpc(reviewService, reviewProducer))

	return &grpcServerObject{
		server: grpcServer,
		cfg:    cfg,
		db:     db,
		log:    log,
	}
}

func (s *grpcServerObject) Start() {
	port := ":50053"
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
