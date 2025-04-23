package dto

type RecipeRequest struct {
	AuthorID    string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
