package repository

import (
	"encoding/json"
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"log"
)

type RecipeRepo struct {
	db interfaces.Database
}

func NewRecipeRepo(db interfaces.Database) *RecipeRepo {
	return &RecipeRepo{
		db: db,
	}
}

func (r *RecipeRepo) RecipeCreateRepo(recipe model.Recipe) (string, error) {
	const op = "handler.recipe.RecipeCreate"

	photosJSON, err := json.Marshal(recipe.Photos)
	if err != nil {
		log.Printf("%s: failed to marshal photos: %v", op, err)
		return "", fmt.Errorf("failed to marshal photos: %w", err)
	}

	recipeEntity := model.RecipeEntity{
		ID:            recipe.ID,
		AuthorID:      recipe.AuthorID,
		Title:         recipe.Title,
		Description:   recipe.Description,
		Photos:        photosJSON,
		AverageRating: float32(recipe.AverageRating),
		CreatedAt:     recipe.CreatedAt,
		UpdatedAt:     recipe.UpdatedAt,
	}

	if err = r.db.GetDB().Create(&recipeEntity).Error; err != nil {
		log.Printf("%s: failed to create recipe: %v", op, err)
		return "", fmt.Errorf("failed to create recipe: %w", err)
	}

	idStr := recipeEntity.ID.String()
	return idStr, nil
}

func (r *RecipeRepo) RecipeUpdateRepo() {}

func (r *RecipeRepo) RecipeListRepo() ([]*model.Recipe, error) {
	const op = "recipe.repository.RecipeListRepo"

	var recipeList []*model.RecipeEntity

	result := r.db.GetDB().Find(&recipeList)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find recipe list: %w", result.Error)
	}
	var recipes []*model.Recipe
	for _, entity := range recipeList {
		recipe := &model.Recipe{
			ID:          entity.ID,
			AuthorID:    entity.AuthorID,
			Title:       entity.Title,
			Description: entity.Description,
			CreatedAt:   entity.CreatedAt,
			UpdatedAt:   entity.UpdatedAt,
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *RecipeRepo) RecipeDeleteRepo() {}
