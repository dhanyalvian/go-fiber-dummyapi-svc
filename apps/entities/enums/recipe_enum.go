// - apps/entities/enums/recipe_enum.go

package enums

type (
	RecipeMealType   string
	RecipeDifficulty string
	RecipeUnit       string
)

const (
	RecipeMealTypeBreakfast  RecipeMealType = "Breakfast"
	RecipeMealTypeLunch      RecipeMealType = "Lunch"
	RecipeMealTypeDinner     RecipeMealType = "Dinner"
	RecipeMealTypeSnack      RecipeMealType = "Snack"
	RecipeMealTypeDessert    RecipeMealType = "Dessert"
	RecipeMealTypeBeverage   RecipeMealType = "Beverage"
	RecipeMealTypeAppetizer  RecipeMealType = "Appetizer"
	RecipeMealTypeSideDish   RecipeMealType = "Side Dish"
	RecipeMealTypeMainCourse RecipeMealType = "Main Course"
	RecipeMealTypeSoup       RecipeMealType = "Soup"
	RecipeMealTypeSalad      RecipeMealType = "Salad"
	RecipeMealTypeCondiment  RecipeMealType = "Condiment"
	RecipeMealTypeBakery     RecipeMealType = "Bakery"
	RecipeMealTypeOther      RecipeMealType = "Other"
)

const (
	RecipeDifficultyEasy   RecipeDifficulty = "Easy"
	RecipeDifficultyMedium RecipeDifficulty = "Medium"
	RecipeDifficultyHard   RecipeDifficulty = "Hard"
)

const (
	RecipeUnitGram       RecipeUnit = "gram"
	RecipeUnitKg         RecipeUnit = "kg"
	RecipeUnitMl         RecipeUnit = "ml"
	RecipeUnitLiter      RecipeUnit = "liter"
	RecipeUnitPcs        RecipeUnit = "pcs"
	RecipeUnitTbsp       RecipeUnit = "tbsp"
	RecipeUnitTsp        RecipeUnit = "tsp"
	RecipeUnitKiloGram   RecipeUnit = "kilogram"
	RecipeUnitMilliliter RecipeUnit = "milliliter"
)
