package server

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/user"
	"github.com/TeslaMode1X/advProg2Final/user/config"
	redis "github.com/TeslaMode1X/advProg2Final/user/internal/infrastructure/cache"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/repository"
	"github.com/TeslaMode1X/advProg2Final/user/internal/service"
	grpcService "github.com/TeslaMode1X/advProg2Final/user/internal/service/grpc"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/nats"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/nats/producer"
	redisconn "github.com/TeslaMode1X/advProg2Final/user/pkg/redis"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type grpcServerObject struct {
	server     *grpc.Server
	cfg        *config.Config
	db         interfaces.Database
	log        *log.Logger
	natsClient *nats.Client
}

func NewGrpcServer(conf *config.Config, db interfaces.Database, log *log.Logger) interfaces.Server {
	ctx := context.Background()

	// NATS connection
	natsClient, err := nats.NewClient(context.Background(), []string{"nats_service:4222"}, "", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	// Sending objects via NATS
	userProducer := producer.NewUserProducer(natsClient)

	// REDIS connection
	log.Println("Attempting to connect to Redis...")
	redisClient, err := redisconn.NewClient(ctx, redisconn.GetRedisConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis!")

	// Test Redis connection with PING
	pingErr := redisClient.Ping(ctx)
	if pingErr != nil {
		log.Fatalf("Redis PING failed: %v", pingErr)
	}
	log.Println("Redis PING successful - connection is working!")

	// REDIS cache
	clientRedisCache := redis.NewClient(redisClient, 10*time.Hour)
	log.Println("Redis cache client initialized with 10 hour TTL")

	userRepository := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepository, clientRedisCache)

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, grpcService.NewUserServiceGrpc(userService, userProducer))

	return &grpcServerObject{
		grpcServer,
		conf,
		db,
		log,
		natsClient,
	}
}

func (s *grpcServerObject) Start() {
	port := ":50051"
	if s.cfg.Server.Port != "" {
		port = ":" + s.cfg.Server.Port
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		s.log.Fatalf("Failed to listen on %s: %v", port, err)
	}

	s.log.Printf("Starting user gRPC server on %s", port)
	if err = s.server.Serve(lis); err != nil {
		s.log.Fatalf("Failed to serve: %v", err)
	}
}
