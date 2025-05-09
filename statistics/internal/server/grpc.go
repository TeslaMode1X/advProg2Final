package server

import (
	"github.com/TeslaMode1X/advProg2Final/proto/gen/statistics"
	"github.com/TeslaMode1X/advProg2Final/statistics/config"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/service"
	grpcService "github.com/TeslaMode1X/advProg2Final/statistics/internal/service/grpc"
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
	statisticsRepository := repository.NewStatisticsRepository(db)
	statisticsService := service.NewStatisticsService(statisticsRepository)

	grpcServer := grpc.NewServer()

	statistics.RegisterStatisticsServiceServer(grpcServer, grpcService.NewStatisticsServiceGrpc(statisticsService))

	return &grpcServerObject{
		server: grpc.NewServer(),
		cfg:    cfg,
		db:     db,
		log:    log,
	}
}

func (s *grpcServerObject) Start() {
	port := ":50054"
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
