package models

import (
	"fmt"
	"go-fiber-dummyapi-svc/apps/dto"
	"go-fiber-dummyapi-svc/apps/entities"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListRecipe(
	c *fiber.Ctx,
	ts *typesense.Client,
	filters *dto.ReqListRecipeFilter,
) (*api.SearchResult, error) {
	queryBy := "name,cuisine,tags,mealType"
	filterBy := getRecipeFilters(filters)
	sortBy := []string{"name:asc"}

	return GetList(c, ts, entities.Recipe{}.ColletionName(), queryBy, filterBy, sortBy)
}

func DetailRecipe(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Recipe{}.ColletionName(), id)
}

func ListRecipeFilters(
	c *fiber.Ctx,
	ts *typesense.Client,
) (entities.RespRecipeFilters, error) {
	result := entities.RespRecipeFilters{
		Cuisine:    []entities.RespFilterOption{},
		Difficulty: []entities.RespFilterOption{},
		MealType:   []entities.RespFilterOption{},
	}

	facets, err := GetFacets(
		c,
		ts,
		entities.Recipe{}.ColletionName(),
		[]string{"cuisine", "difficulty", "mealType"},
	)
	if err != nil {
		return result, err
	}

	return mapRecipeFacets(facets, result), nil
}

func mapRecipeFacets(
	facets *api.SearchResult,
	result entities.RespRecipeFilters,
) entities.RespRecipeFilters {
	if facets == nil || facets.FacetCounts == nil {
		return result
	}

	for _, facet := range *facets.FacetCounts {
		if facet.FieldName == nil {
			continue
		}

		options := []entities.RespFilterOption{}
		if facet.Counts != nil {
			for _, count := range *facet.Counts {
				if count.Value == nil {
					continue
				}

				options = append(
					options,
					entities.RespFilterOption{
						Label: *count.Value,
						Value: *count.Value,
					},
				)
			}
		}

		switch *facet.FieldName {
		case "cuisine":
			result.Cuisine = options

		case "difficulty":
			result.Difficulty = options

		case "mealType":
			result.MealType = options
		}
	}

	return result
}

func getRecipeFilters(filters *dto.ReqListRecipeFilter) []string {
	var result []string

	if filters == nil {
		return result
	}

	//- filter by Cuisine
	result = getRecipeFilterCuisine(filters, result)

	//- filter by Difficulty
	result = getRecipeFilterDifficulty(filters, result)

	//- filter by Meal Type
	result = getRecipeFilterMealType(filters, result)

	return result
}

func getRecipeFilterCuisine(filters *dto.ReqListRecipeFilter, result []string) []string {
	if filters.Cuisine != nil && len(*filters.Cuisine) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"cuisine:=[%s]",
				strings.Join(*filters.Cuisine, ","),
			),
		)
	}

	return result
}

func getRecipeFilterDifficulty(filters *dto.ReqListRecipeFilter, result []string) []string {
	if filters.Difficulty != nil && len(*filters.Difficulty) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"difficulty:=[%s]",
				strings.Join(*filters.Difficulty, ","),
			),
		)
	}

	return result
}

func getRecipeFilterMealType(filters *dto.ReqListRecipeFilter, result []string) []string {
	if filters.MealType != nil && len(*filters.MealType) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"mealType:=[%s]",
				strings.Join(*filters.MealType, ","),
			),
		)
	}

	return result
}
