package producer

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/TeslaMode1X/advProg2Final/proto/gen/recipe"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"github.com/TeslaMode1X/advProg2Final/recipe/pkg/nats"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

const (
	Subject     = "recipes.recipe"
	PushTimeout = time.Second * 30
)

type RecipeProducer struct {
	client *nats.Client
}

func NewRecipeProducer(client *nats.Client) *RecipeProducer {
	return &RecipeProducer{
		client: client,
	}
}

func (p *RecipeProducer) PublishRecipeCreated(ctx context.Context, recipe model.RecipeNats) error {
	data, err := json.Marshal(recipe)
	if err != nil {
		return fmt.Errorf("failed to marshal recipe: %w", err)
	}

	err = p.client.Conn.Publish(Subject, data)
	if err != nil {
		return fmt.Errorf("failed to publish recipe created event: %w", err)
	}

	log.Printf("Published recipe created event: %s", recipe.AuthorID)
	return nil
}

func (p *RecipeProducer) Push(ctx context.Context, recipe model.RecipeNats) error {
	ctx, cancel := context.WithTimeout(ctx, PushTimeout)
	defer cancel()

	recipePb := &pb.RecipeNats{
		AuthorId: recipe.AuthorID.String(),
	}
	data, err := proto.Marshal(recipePb)
	if err != nil {
		return fmt.Errorf("proto.Marshal: %w", err)
	}

	err = p.client.Conn.Publish(Subject, data)
	if err != nil {
		return fmt.Errorf("p.recipe.Conn.Publish: %w", err)
	}
	log.Println("recipe is pushed:", recipe.AuthorID)

	return nil
}
