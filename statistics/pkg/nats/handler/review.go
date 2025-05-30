package handler

import (
	"context"
	"encoding/json"
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/model"
	reviewModel "github.com/TeslaMode1X/advProg2Final/statistics/pkg/nats/model/review"
	"github.com/gofrs/uuid"
	"github.com/nats-io/nats.go"

	"log"
)

type ReviewHandler struct {
	service interfaces.StatisticsService
}

func NewReviewHandler(service interfaces.StatisticsService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) HandlerReview(ctx context.Context, msg *nats.Msg) error {
	log.Printf("Handler called with subject: %s and data: %s", msg.Subject, string(msg.Data))

	var review reviewModel.Nats
	if err := json.Unmarshal(msg.Data, &review); err != nil {
		log.Printf("Failed to unmarshal review data: %v", err)
		return fmt.Errorf("failed to unmarshal review data: %w", err)
	}

	log.Printf("Received user registration event: %s", review.AuthorId)

	authorUUID, _ := uuid.FromString(review.AuthorId)
	recipeUUID, _ := uuid.FromString(review.RecipeId)

	modelNeed := &model.RecipeReviewStatistics{
		ID:          authorUUID,
		RecipeID:    recipeUUID,
		TotalRating: float32(review.Rating),
	}

	if err := h.service.AddNewReview(*modelNeed); err != nil {
		log.Printf("Failed to update user statistics: %v", err)
		return fmt.Errorf("failed to update user statistics: %w", err)
	}

	return nil
}
