package grpc

import (
	"context"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/statistics"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatisticsServiceGrpc struct {
	statistics.UnimplementedStatisticsServiceServer
	statisticsService interfaces.StatisticsService
}

func NewStatisticsServiceGrpc(statisticsService interfaces.StatisticsService) *StatisticsServiceGrpc {
	return &StatisticsServiceGrpc{
		statisticsService: statisticsService,
	}
}

func (s *StatisticsServiceGrpc) GetUsersStatistics(ctx context.Context, req statistics.Empty) (*statistics.StatisticsUserGetResponse, error) {
	userStatisticsRequest, err := s.statisticsService.GetUsersStatisticsService()
	if err != nil {
		return nil, err
	}

	userIDstr := userStatisticsRequest.ID.String()

	userStatisticsReturn := &statistics.StatisticsUserGetResponse{
		Id:            userIDstr,
		TotalUsers:    int32(userStatisticsRequest.TotalUsers),
		LastUpdatedAt: timestamppb.New(userStatisticsRequest.LastUpdatedAt),
	}

	return userStatisticsReturn, nil
}

func (s *StatisticsServiceGrpc) GetStatisticsRecipes(ctx context.Context, req statistics.Empty) (*statistics.StatisticsRecipesResponse, error) {
	statisticsRecipe, err := s.statisticsService.GetRecipesStatisticsService()
	if err != nil {
		return nil, err
	}

	var statisticsRecipesReturn []*statistics.StatisticsRecipeByIDResponse

	for _, recipe := range statisticsRecipe {
		var statisticsRecipeReturn statistics.StatisticsRecipeByIDResponse
		statisticsRecipeReturn.Id = recipe.ID.String()
		statisticsRecipeReturn.RecipeId = recipe.RecipeID.String()
		statisticsRecipeReturn.TotalReviews = int32(recipe.TotalReviews)
		statisticsRecipeReturn.TotalRating = recipe.TotalRating
		statisticsRecipeReturn.AverageRating = float32(recipe.AverageRating)
		statisticsRecipeReturn.LastUpdatedAt = timestamppb.New(recipe.LastUpdatedAt)

		statisticsRecipesReturn = append(statisticsRecipesReturn, &statisticsRecipeReturn)
	}

	return &statistics.StatisticsRecipesResponse{
		Statistics: statisticsRecipesReturn,
	}, nil
}

func (s *StatisticsServiceGrpc) GetStatisticsRecipeByID(ctx context.Context, req statistics.StatisticsRecipeByIDRequest) (*statistics.StatisticsRecipeByIDResponse, error) {
	recipeStatistics, err := s.statisticsService.GetRecipeStatByIDService(req.Id)
	if err != nil {
		return nil, err
	}

	var statisticsRecipeReturn statistics.StatisticsRecipeByIDResponse
	statisticsRecipeReturn.Id = recipeStatistics.ID.String()
	statisticsRecipeReturn.RecipeId = recipeStatistics.RecipeID.String()
	statisticsRecipeReturn.TotalReviews = int32(recipeStatistics.TotalReviews)
	statisticsRecipeReturn.TotalRating = recipeStatistics.TotalRating
	statisticsRecipeReturn.AverageRating = recipeStatistics.AverageRating
	statisticsRecipeReturn.LastUpdatedAt = timestamppb.New(recipeStatistics.LastUpdatedAt)

	return &statisticsRecipeReturn, nil
}
