package grpc

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/review"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/gofrs/uuid"
)

type ReviewServerGrpc struct {
	review.UnimplementedReviewServiceServer
	reviewService interfaces.ReviewService
}

func NewReviewServerGrpc(s interfaces.ReviewService) *ReviewServerGrpc {
	return &ReviewServerGrpc{
		reviewService: s,
	}
}

func (r *ReviewServerGrpc) ReviewCreate(ctx context.Context, req *review.ReviewCreateRequest) (*review.ReviewCreateResponse, error) {
	userIDString, _ := uuid.FromString(req.UserId)
	recipeIDString, _ := uuid.FromString(req.RecipeId)

	reviewObject := &model.Review{
		UserID:   userIDString,
		RecipeID: recipeIDString,
		Comment:  req.Comment,
		Rating:   req.Rating,
	}

	id, err := r.reviewService.ReviewCreateService(reviewObject)
	if err != nil {
		return nil, err
	}

	return &review.ReviewCreateResponse{
		Id: id,
	}, nil
}

func (r *ReviewServerGrpc) ReviewGetList(ctx context.Context, req *review.Empty) (*review.ReviewGetListResponse, error) {
	objectList, err := r.reviewService.ReviewListService()
	if err != nil {
		return nil, err
	}

	var returnObject review.ReviewGetListResponse
	for _, object := range objectList {
		obj := &review.Review{
			Id:       object.ID.String(),
			UserId:   object.UserID.String(),
			RecipeId: object.RecipeID.String(),
			Comment:  object.Comment,
			Rating:   object.Rating,
		}
		returnObject.Reviews = append(returnObject.Reviews, obj)
	}

	return &returnObject, nil
}

func (r *ReviewServerGrpc) ReviewGetById(ctx context.Context, req *review.ReviewGetByIdRequest) (*review.ReviewGetByIdResponse, error) {
	object, err := r.reviewService.ReviewByIDService(req.Id)
	if err != nil {
		return nil, err
	}

	obj := &review.Review{
		Id:       object.ID.String(),
		UserId:   object.UserID.String(),
		RecipeId: object.RecipeID.String(),
		Comment:  object.Comment,
		Rating:   object.Rating,
	}

	return &review.ReviewGetByIdResponse{
		Review: obj,
	}, nil
}

func (r *ReviewServerGrpc) ReviewUpdate(ctx context.Context, req *review.ReviewUpdateRequest) (*review.ReviewUpdateResponse, error) {
	modelID, _ := uuid.FromString(req.Id)
	userID, _ := uuid.FromString(req.UserId)
	recipeID, _ := uuid.FromString(req.RecipeId)

	modelObject := &model.Review{
		ID:       modelID,
		UserID:   userID,
		RecipeID: recipeID,
		Comment:  req.Comment,
		Rating:   req.Rating,
	}

	err := r.reviewService.ReviewUpdateService(modelObject)
	if err != nil {
		return nil, err
	}

	reviewObject := &review.Review{
		Id:       req.Id,
		UserId:   req.UserId,
		RecipeId: req.RecipeId,
		Comment:  req.Comment,
		Rating:   req.Rating,
	}

	return &review.ReviewUpdateResponse{
		Review: reviewObject,
	}, nil
}

func (r *ReviewServerGrpc) ReviewDelete(ctx context.Context, req *review.ReviewDeleteRequest) (*review.ReviewDeleteResponse, error) {
	err := r.reviewService.ReviewDeleteService(req.Id)
	if err != nil {
		return nil, err
	}

	return &review.ReviewDeleteResponse{
		Status: "deleted",
	}, nil
}
