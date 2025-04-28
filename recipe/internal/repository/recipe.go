package repository

import (
	"encoding/json"
	"fmt"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/repository/dao"
	"github.com/gofrs/uuid"
	"log"
	"os"
	"time"
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
	const op = "repository.recipe.RecipeCreateRepo"

	photosJSON, err := json.Marshal(recipe.Photos)
	if err != nil {
		log.Printf("%s: failed to marshal photos: %v", op, err)
		return "", fmt.Errorf("failed to marshal photos: %w", err)
	}

	recipeEntity := dao.RecipeEntity{
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

func (r *RecipeRepo) RecipeByIDRepo(id string) (*dao.RecipeEntity, error) {
	const op = "repository.recipe.RecipeByIDRepo"

	var recipeObject dao.RecipeEntity

	result := r.db.GetDB().Where("id = ?", id).First(&recipeObject)
	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, result.Error)
	}

	return &recipeObject, nil
}

func (r *RecipeRepo) RecipeUpdateRepo(recipe model.Recipe) error {
	const op = "repository.recipe.RecipeUpdateRepo"

	var existingRecipe dao.RecipeEntity
	if err := r.db.GetDB().Where("id = ?", recipe.ID).First(&existingRecipe).Error; err != nil {
		return fmt.Errorf("%s: failed to find existing recipe: %w", op, err)
	}

	if recipe.Photos != nil {
		if len(existingRecipe.Photos) > 0 {
			var existingPhotos []string
			if err := json.Unmarshal(existingRecipe.Photos, &existingPhotos); err != nil {
				return fmt.Errorf("%s: failed to unmarshal existing photos: %w", op, err)
			}

			// Удаляем старые файлы
			for _, path := range existingPhotos {
				fullPath := fmt.Sprintf("../%s", path)
				err := os.Remove(fullPath)
				if err != nil && !os.IsNotExist(err) {
					log.Printf("Failed to delete photo %s: %v", fullPath, err)
				}
			}
		}

		photosJSON, err := json.Marshal(recipe.Photos)
		if err != nil {
			return fmt.Errorf("%s: failed to marshal new photos: %w", op, err)
		}

		result := r.db.GetDB().Model(&dao.RecipeEntity{}).
			Where("id = ?", recipe.ID).
			Updates(map[string]interface{}{
				"title":       recipe.Title,
				"description": recipe.Description,
				"photos":      photosJSON,
				"updated_at":  time.Now(),
			})

		if result.Error != nil {
			return fmt.Errorf("%s: failed to update recipe: %w", op, result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("%s: recipe with ID %s not found", op, recipe.ID)
		}
	} else {
		result := r.db.GetDB().Model(&dao.RecipeEntity{}).
			Where("id = ?", recipe.ID).
			Updates(map[string]interface{}{
				"title":       recipe.Title,
				"description": recipe.Description,
				"updated_at":  time.Now(),
			})

		if result.Error != nil {
			return fmt.Errorf("%s: failed to update recipe: %w", op, result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("%s: recipe with ID %s not found", op, recipe.ID)
		}
	}

	return nil
}

func (r *RecipeRepo) RecipeListRepo() ([]*model.Recipe, error) {
	const op = "repository.recipe.RecipeListRepo"

	var recipeList []*dao.RecipeEntity

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

func (r *RecipeRepo) RecipeDeleteRepo(id string) error {
	const op = "repository.recipe.RecipeDeleteRepo"

	entity := dao.RecipeEntity{}
	parsedID, err := uuid.FromString(id)
	if err != nil {
		return fmt.Errorf("%s: invalid UUID format: %w", op, err)
	}
	entity.ID = parsedID

	result := r.db.GetDB().Delete(&entity)
	if result.Error != nil {
		return fmt.Errorf("%s: failed to delete entity: %w", op, result.Error)
	}

	return nil
}

func (r *RecipeRepo) RecipeUserCheck(recipeID, userID string) (bool, error) {
	const op = "repository.recipe.RecipeUserCheck"

	var count int64
	result := r.db.GetDB().Model(&dao.RecipeEntity{}).
		Where("id = ? AND author_id = ?", recipeID, userID).
		Count(&count)

	if result.Error != nil {
		return false, fmt.Errorf("%s: failed to check user's recipe: %w", op, result.Error)
	}

	return count > 0, nil
}
