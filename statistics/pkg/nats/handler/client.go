package handler

import (
	"context"
	"encoding/json"
	"fmt"
	userModel "github.com/TeslaMode1X/advProg2Final/statistics/pkg/nats/model/user"

	"log"

	"github.com/nats-io/nats.go"

	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
)

type UserHandler struct {
	service interfaces.StatisticsService
}

func NewUserHandler(service interfaces.StatisticsService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Handler(ctx context.Context, msg *nats.Msg) error {
	log.Printf("Handler called with subject: %s and data: %s", msg.Subject, string(msg.Data))

	var user userModel.Nats
	if err := json.Unmarshal(msg.Data, &user); err != nil {
		log.Printf("Failed to unmarshal user data: %v", err)
		return fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	log.Printf("Received user registration event: %s", user.Email)

	if err := h.service.AddNewUserCounter(); err != nil {
		log.Printf("Failed to update user statistics: %v", err)
		return fmt.Errorf("failed to update user statistics: %w", err)
	}

	return nil
}
