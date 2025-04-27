package dto

import (
	"github.com/TeslaMode1X/advProg2Final/recipe/pkg/errors"
)

type RecipeRequest struct {
	AuthorID    string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *RecipeRequest) Validate() error {
	if r.AuthorID == "" {
		return errors.ErrorAuthorIDIsRequired
	}

	if r.Title == "" {
		return errors.ErrorTitleIsRequired
	}

	if r.Description == "" {
		return errors.ErrorDescriptionIsRequired
	}

	return nil
}

type RecipeResponse struct {
	AuthorID    string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RecipeUpdateRequest struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}
