package service

import (
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
)

type ReviewService struct {
	userRepo interfaces.ReviewRepository
}

func NewReviewService(userRepo interfaces.ReviewRepository) *ReviewService {
	return &ReviewService{
		userRepo: userRepo,
	}
}

func (s *ReviewService) ReviewCreateService(model *model.Review) (string, error) {
	const op = "service.review.ReviewCreateService"

	id, err := s.userRepo.ReviewCreateRepo(model)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *ReviewService) ReviewListService() ([]*dao.ReviewEntity, error) {
	const op = "service.review.ReviewListService"

	list, err := s.userRepo.ReviewListRepo()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return list, nil
}
