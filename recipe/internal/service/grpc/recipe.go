package grpc

import (
	"context"
	"encoding/json"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/recipe"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecipeServerGrpc struct {
	recipe.UnimplementedRecipeServiceServer
	recipeService interfaces.RecipeService
}

func NewRecipeServerGrpc(s interfaces.RecipeService) *RecipeServerGrpc {
	return &RecipeServerGrpc{
		recipeService: s,
	}
}

func (r *RecipeServerGrpc) RecipeList(ctx context.Context, req *recipe.Empty) (*recipe.RecipeListResponse, error) {
	recipeObjects, err := r.recipeService.RecipeListService()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list recipes: %s", err)
	}

	var recipes []*recipe.Recipe
	for _, obj := range recipeObjects {
		recipes = append(recipes, &recipe.Recipe{
			Id:          obj.ID.String(),
			Title:       obj.Title,
			Description: obj.Description,
			Photos:      obj.Photos,
		})
	}

	return &recipe.RecipeListResponse{
		Recipes: recipes,
	}, nil
}

func (r *RecipeServerGrpc) RecipeByID(ctx context.Context, req *recipe.RecipeByIDRequest) (*recipe.RecipeByIDResponse, error) {
	recipeObject, err := r.recipeService.RecipeByIDService(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get the recipe: %s", err)
	}

	var photos []string
	if err = json.Unmarshal(recipeObject.Photos, &photos); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal the recipe: %s", err)
	}

	return &recipe.RecipeByIDResponse{
		Recipe: &recipe.Recipe{
			Id:          recipeObject.ID.String(),
			Title:       recipeObject.Title,
			Description: recipeObject.Description,
			Photos:      photos,
		},
	}, nil
}

func (r *RecipeServerGrpc) RecipeCreate(ctx context.Context, req *recipe.RecipeCreateRequest) (*recipe.RecipeCreateResponse, error) {
	recipeUUID, _ := uuid.NewV4()

	authorIDString, _ := uuid.FromString(req.AuthorId)

	recipeCreate := &model.Recipe{
		ID:          recipeUUID,
		AuthorID:    authorIDString,
		Title:       req.Title,
		Description: req.Description,
		Photos:      req.Photos,
	}

	id, err := r.recipeService.RecipeCreateService(*recipeCreate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create the reciep: %s", err)
	}

	return &recipe.RecipeCreateResponse{
		Id: id,
	}, nil
}

func (r *RecipeServerGrpc) RecipeUpdate(ctx context.Context, req *recipe.RecipeUpdateRequest) (*recipe.RecipeUpdateResponse, error) {
	recipeUUID, _ := uuid.FromString(req.Id)

	recipeUpdateObject := &model.Recipe{
		ID:          recipeUUID,
		Title:       req.Title,
		Description: req.Description,
		Photos:      req.Photos,
	}

	err := r.recipeService.RecipeUpdateService(*recipeUpdateObject)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update the recipe: %s", err)
	}

	return &recipe.RecipeUpdateResponse{
		Id: recipeUUID.String(),
	}, nil
}

func (r *RecipeServerGrpc) RecipeDelete(ctx context.Context, req *recipe.RecipeDeleteRequest) (*recipe.RecipeDeleteResponse, error) {
	err := r.recipeService.RecipeDeleteService(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete the recipe: %s", err)
	}

	return &recipe.RecipeDeleteResponse{
		Id: req.Id,
	}, nil
}
