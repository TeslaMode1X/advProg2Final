package service

import (
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
)

type RecipeService struct {
	recipeRepo interfaces.RecipeRepository
}

func NewRecipeService(recipeRepo interfaces.RecipeRepository) *RecipeService {
	return &RecipeService{
		recipeRepo: recipeRepo,
	}
}

func (s *RecipeService) RecipeCreateService(recipe model.Recipe) (string, error) {
	const op = "handler.recipe.CreateRecipe"

	id, err := s.recipeRepo.RecipeCreateRepo(recipe)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *RecipeService) RecipeListService() ([]*model.Recipe, error) {
	const op = "recipe.service.RecipeListService"

	recipes, err := s.recipeRepo.RecipeListRepo()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return recipes, nil
}

func (s *RecipeService) RecipeUpdateService() {
	//TODO implement me
	panic("implement me")
}

func (s *RecipeService) RecipeDeleteService() {
	//TODO implement me
	panic("implement me")
}
