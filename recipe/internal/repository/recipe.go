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

	if err := r.db.GetDB().Create(&recipeEntity).Error; err != nil {
		log.Printf("%s: failed to create recipe: %v", op, err)
		return "", fmt.Errorf("failed to create recipe: %w", err)
	}

	idStr := recipeEntity.ID.String()
	return idStr, nil
}

func (r *RecipeRepo) RecipeUpdateRepo() {}

func (r *RecipeRepo) RecipeListRepo() {}

func (r *RecipeRepo) RecipeDeleteRepo() {}
