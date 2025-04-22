package interfaces

type RecipeRepository interface {
	RecipeListRepo()
	RecipeCreateRepo()
	RecipeUpdateRepo()
	RecipeDeleteRepo()
}
