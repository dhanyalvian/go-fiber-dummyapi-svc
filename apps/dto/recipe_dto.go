//- apps/dto/recipe_dto.go

package dto

type ReqListRecipeFilter struct {
	Cuisine    *[]string `json:"cuisine"`
	Difficulty *[]string `json:"difficulty"`
	MealType   *[]string `json:"mealType"`
}
