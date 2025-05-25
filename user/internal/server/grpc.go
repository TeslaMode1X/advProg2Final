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
	server             *grpc.Server
	cfg                *config.Config
	db                 interfaces.Database
	log                *log.Logger
	natsClient         *nats.Client
	cacheRefreshCancel context.CancelFunc
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

	server := &grpcServerObject{
		server:     grpcServer,
		cfg:        conf,
		db:         db,
		log:        log,
		natsClient: natsClient,
	}

	server.startCacheRefreshJob(ctx, userService, 10*time.Hour)

	user.RegisterUserServiceServer(grpcServer, grpcService.NewUserServiceGrpc(userService, userProducer))

	return server
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

func (s *grpcServerObject) startCacheRefreshJob(ctx context.Context, userService *service.UserService, refreshInterval time.Duration) {
	refreshCtx, cancel := context.WithCancel(ctx)
	s.cacheRefreshCancel = cancel

	go func() {
		ticker := time.NewTicker(refreshInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println("Running scheduled cache refresh...")
				err := userService.RefreshCache()
				if err != nil {
					log.Printf("Scheduled cache refresh failed: %v", err)
				} else {
					log.Println("Scheduled cache refresh completed successfully")
				}
			case <-refreshCtx.Done():
				log.Println("Cache refresh job terminated")
				return
			}
		}
	}()

	log.Printf("Background cache refresh job started with %v interval", refreshInterval)
}
