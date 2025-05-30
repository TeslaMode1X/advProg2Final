package server

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/statistics"
	"github.com/TeslaMode1X/advProg2Final/statistics/config"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/service"
	grpcService "github.com/TeslaMode1X/advProg2Final/statistics/internal/service/grpc"
	"github.com/TeslaMode1X/advProg2Final/statistics/pkg/nats"
	"github.com/TeslaMode1X/advProg2Final/statistics/pkg/nats/consumer"
	natsHandler "github.com/TeslaMode1X/advProg2Final/statistics/pkg/nats/handler"
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
	pubSub     *consumer.PubSub
}

func NewGrpcServer(cfg *config.Config, db interfaces.Database, log *log.Logger) *grpcServerObject {
	statisticsRepository := repository.NewStatisticsRepository(db)
	statisticsService := service.NewStatisticsService(statisticsRepository)

	grpcServer := grpc.NewServer()

	s := &grpcServerObject{
		server: grpcServer,
		cfg:    cfg,
		db:     db,
		log:    log,
	}

	var err error
	s.natsClient, err = nats.NewClient(context.Background(), []string{"nats_service:4222"}, "", true)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	log.Println("NATS connection status is", s.natsClient.Conn.Status().String())

	s.pubSub = consumer.NewPubSub(s.natsClient)

	userHandler := natsHandler.NewUserHandler(statisticsService)
	s.pubSub.Subscribe(consumer.PubSubSubscriptionConfig{
		Subject: "users.user",
		Handler: userHandler.Handler,
	})

	reviewHandler := natsHandler.NewReviewHandler(statisticsService)
	s.pubSub.Subscribe(consumer.PubSubSubscriptionConfig{
		Subject: "reviews.review",
		Handler: reviewHandler.HandlerReview,
	})

	statistics.RegisterStatisticsServiceServer(grpcServer, grpcService.NewStatisticsServiceGrpc(statisticsService))
	
	return s
}

func (s *grpcServerObject) Start() {
	port := ":50054"
	if s.cfg.Server.Port != "" {
		port = ":" + s.cfg.Server.Port
	}

	errCh := make(chan error, 1)
	s.pubSub.Start(context.Background(), errCh)

	go func() {
		for err := range errCh {
			s.log.Printf("NATS error: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		s.log.Fatalf("Failed to listen on %s: %v", port, err)
	}

	s.log.Printf("Starting recipe gRPC server on %s", port)
	if err = s.server.Serve(lis); err != nil {
		s.log.Fatalf("Failed to serve: %v", err)
	}
}
