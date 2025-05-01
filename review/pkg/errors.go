package pkg

import "errors"

var (
	ErrRecipeIdEmpty = errors.New("missing recipeID")
	ErrUserIdEmpty   = errors.New("missing user ID")
	ErrInvalidRating = errors.New("invalid rating")
	ErrCommentEmpty  = errors.New("missing comment")
)
