package producer

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/TeslaMode1X/advProg2Final/proto/gen/review"
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/pkg/nats"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

const (
	Subject     = "reviews.review"
	PushTimeout = time.Second * 30
)

type ReviewProducer struct {
	client *nats.Client
}

func NewReviewProducer(client *nats.Client) *ReviewProducer {
	return &ReviewProducer{
		client: client,
	}
}

func (p *ReviewProducer) PublishReviewCreated(ctx context.Context, review model.ReviewNats) error {
	data, err := json.Marshal(review)
	if err != nil {
		return fmt.Errorf("failed to marshal review: %w", err)
	}

	err = p.client.Conn.Publish(Subject, data)
	if err != nil {
		return fmt.Errorf("failed to publish review created event: %w", err)
	}

	log.Printf("Published review created event: %+v", review)
	return nil
}

func (p *ReviewProducer) Push(ctx context.Context, review model.ReviewNats) error {
	ctx, cancel := context.WithTimeout(ctx, PushTimeout)
	defer cancel()

	reviewPb := &pb.ReviewNats{
		AuthorId: review.AuthorID.String(),
		RecipeId: review.RecipeID.String(),
		Rating:   review.Rating,
	}
	data, err := proto.Marshal(reviewPb)
	if err != nil {
		return fmt.Errorf("proto.Marshal: %w", err)
	}

	err = p.client.Conn.Publish(Subject, data)
	if err != nil {
		return fmt.Errorf("p.review.Conn.Publish: %w", err)
	}
	log.Println("review is pushed:", review)

	return nil
}
