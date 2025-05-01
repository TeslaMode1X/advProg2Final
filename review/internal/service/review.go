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

func (s *ReviewService) ReviewByIDService(id string) (*dao.ReviewEntity, error) {
	const op = "service.review.ReviewByIDService"

	object, err := s.userRepo.ReviewByIDRepo(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return object, nil
}

func (s *ReviewService) ReviewUpdateService(modelObject *model.Review) error {
	const op = "service.review.ReviewUpdateService"

	err := s.userRepo.ReviewUpdateRepo(modelObject)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *ReviewService) ReviewDeleteService(id string) error {
	const op = "service.review.ReviewDeleteService"

	err := s.userRepo.ReviewDeleteRepo(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
