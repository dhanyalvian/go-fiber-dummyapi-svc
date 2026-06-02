//- apps/entities/recipe_entity.go

package entities

import (
	"go-fiber-dummyapi-svc/apps/entities/enums"
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/typesense/typesense-go/v4/typesense/api"
)

type Recipe struct {
	BaseID

	Name         string                 `json:"name" typesense:"index,sort"`
	Cuisine      string                 `json:"cuisine" typesense:"index,facet"`
	CuisineCode  string                 `json:"cuisineCode" typesense:"index,facet"`
	Description  string                 `json:"description" typesense:"optional"`
	Difficulty   enums.RecipeDifficulty `json:"difficulty" typesense:"index,facet,sort"`
	MealType     []string               `json:"mealType" typesense:"index,facet"`
	Tags         []string               `json:"tags" typesense:"index"`
	PrepTime     int                    `json:"prepTime"` // menit
	CookTime     int                    `json:"cookTime"` // menit
	Servings     int                    `json:"servings"`
	Image        string                 `json:"image" typesense:"optional"`
	Ingredients  []Ingredient           `json:"ingredients"`
	Instructions []Instruction          `json:"instructions"`

	BaseTimestamp
}

type Ingredient struct {
	ID       int64            `json:"id" typesense:"sort"`
	Name     string           `json:"name"`
	Quantity float64          `json:"quantity"`
	Unit     enums.RecipeUnit `json:"unit"`
}

type Instruction struct {
	ID          int64  `json:"id"`
	StepNo      int    `json:"stepNo" typesense:"sort"`
	Description string `json:"description"`
}

type RecipeDoc struct {
	Recipe
}

type RespListRecipe struct {
	BaseID

	Name        string `json:"name"`
	Image       string `json:"image"`
	Cuisine     string `json:"cuisine"`
	CuisineCode string `json:"cuisineCode"`

	Difficulty enums.RecipeDifficulty `json:"difficulty"`
	MealType   []string               `json:"mealType"`
	Tags       []string               `json:"tags"`
}

type RespDetailRecipe struct {
	RespListRecipe

	Description  string        `json:"description"`
	PrepTime     int           `json:"prepTime"` // menit
	CookTime     int           `json:"cookTime"` // menit
	Servings     int           `json:"servings"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
}

func (Recipe) ColletionName() string {
	return GetCollectionName(COLLECTION_RECIPE)
}

func (Recipe) TypesenseSchema() ([]api.Field, *string) {
	return utils.DeriveTypesenseFields[Recipe](), nil
}
