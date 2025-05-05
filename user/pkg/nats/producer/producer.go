package producer

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/TeslaMode1X/advProg2Final/proto/gen/user"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/TeslaMode1X/advProg2Final/user/pkg/nats"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

const (
	Subject     = "users.user"
	PushTimeout = time.Second * 30
)

type UserProducer struct {
	client *nats.Client
}

func NewUserProducer(client *nats.Client) *UserProducer {
	return &UserProducer{
		client: client,
	}
}

func (p *UserProducer) PublishUserCreated(ctx context.Context, user model.UserNats) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	err = p.client.Conn.Publish(Subject, data)
	if err != nil {
		return fmt.Errorf("failed to publish user created event: %w", err)
	}

	log.Printf("Published user created event: %s", user.Email)
	return nil
}

func (p *UserProducer) Push(ctx context.Context, product model.UserNats) error {
	ctx, cancel := context.WithTimeout(ctx, PushTimeout)
	defer cancel()

	userPb := &pb.UserNatsRequest{
		Email: product.Email,
	}
	data, err := proto.Marshal(userPb)
	if err != nil {
		return fmt.Errorf("proto.Marshal: %w", err)
	}

	err = p.client.Conn.Publish(Subject, data)
	if err != nil {
		return fmt.Errorf("p.user.Conn.Publish: %w", err)
	}
	log.Println("user is pushed:", product)

	return nil
}
