package server

import (
	"github.com/TeslaMode1X/advProg2Final/proto/gen/review"
	"github.com/TeslaMode1X/advProg2Final/review/config"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/review/internal/service"
	grpcService "github.com/TeslaMode1X/advProg2Final/review/internal/service/grpc"
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

func NewGrpcServer(cfg *config.Config, db interfaces.Database, log *log.Logger) *grpcServerObject {
	reviewRepository := repository.NewReviewRepo(db)
	reviewService := service.NewReviewService(reviewRepository)

	grpcServer := grpc.NewServer()

	review.RegisterReviewServiceServer(grpcServer, grpcService.NewReviewServerGrpc(reviewService))

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
