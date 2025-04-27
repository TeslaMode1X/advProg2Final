package service

import (
	"encoding/json"
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/repository/dao"
	"log"
	"os"
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
	const op = "service.recipe.CreateRecipe"

	id, err := s.recipeRepo.RecipeCreateRepo(recipe)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *RecipeService) RecipeByIDService(id string) (*dao.RecipeEntity, error) {
	const op = "service.recipe.RecipeByIDService"

	recipeObject, err := s.recipeRepo.RecipeByIDRepo(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return recipeObject, nil
}

func (s *RecipeService) RecipeListService() ([]*model.Recipe, error) {
	const op = "service.recipe.RecipeListService"

	recipes, err := s.recipeRepo.RecipeListRepo()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return recipes, nil
}

func (s *RecipeService) RecipeUpdateService(recipe model.Recipe) error {
	const op = "service.recipe.RecipeUpdateService"

	recipeObject, err := s.recipeRepo.RecipeByIDRepo(recipe.ID.String())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = DeletePhotos(*recipeObject)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.recipeRepo.RecipeUpdateRepo(recipe)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *RecipeService) RecipeDeleteService(id string) error {
	const op = "service.recipe.RecipeDeleteService"

	recipeObject, err := s.recipeRepo.RecipeByIDRepo(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = DeletePhotos(*recipeObject)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.recipeRepo.RecipeDeleteRepo(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func DeletePhotos(recipeObject dao.RecipeEntity) error {
	var photoPaths []string
	if err := json.Unmarshal(recipeObject.Photos, &photoPaths); err != nil {
		return fmt.Errorf("%s: failed to unmarshal photo paths: %w", err)
	}

	for _, path := range photoPaths {
		fullPath := fmt.Sprintf("../%s", path)

		err := os.Remove(fullPath)
		if err != nil && !os.IsNotExist(err) {
			log.Printf("Failed to delete photo %s: %v", fullPath, err)
		}
	}

	return nil
}
