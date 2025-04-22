package interfaces

type RecipeService interface {
	RecipeListService()
	RecipeCreateService()
	RecipeUpdateService()
	RecipeDeleteService()
}
