//- apps/models/product_model.go

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

func ListProduct(
	c *fiber.Ctx,
	ts *typesense.Client,
	filters *dto.ReqListProductFilter,
) (*api.SearchResult, error) {
	queryBy := "name,sku,brand,category,tags"
	filterBy := getProductFilters(filters)
	sortBy := []string{"name:asc"}

	return GetList(c, ts, entities.Product{}.ColletionName(), queryBy, filterBy, sortBy)
}

func DetailProduct(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Product{}.ColletionName(), id)
}

func ListProductFilters(
	c *fiber.Ctx,
	ts *typesense.Client,
) (entities.RespProductFilters, error) {
	result := entities.RespProductFilters{
		Brand:    []entities.RespFilterOption{},
		Category: []entities.RespFilterOption{},
	}

	facets, err := GetFacets(
		c,
		ts,
		entities.Product{}.ColletionName(),
		[]string{"brand", "category"},
	)
	if err != nil {
		return result, err
	}

	return mapProductFacets(facets, result), nil
}

func mapProductFacets(
	facets *api.SearchResult,
	result entities.RespProductFilters,
) entities.RespProductFilters {
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
		case "brand":
			result.Brand = options

		case "category":
			result.Category = options
		}
	}

	return result
}

func getProductFilters(filters *dto.ReqListProductFilter) []string {
	var result []string

	if filters == nil {
		return result
	}

	//- filter by Brand
	result = getProductFilterBrand(filters, result)

	//- filter by Category
	result = getProductFilterCategory(filters, result)

	return result
}

func getProductFilterBrand(filters *dto.ReqListProductFilter, result []string) []string {
	if filters.Brand != nil && len(*filters.Brand) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"brand:=[%s]",
				strings.Join(*filters.Brand, ","),
			),
		)
	}

	return result
}

func getProductFilterCategory(filters *dto.ReqListProductFilter, result []string) []string {
	if filters.Category != nil && len(*filters.Category) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"category:=[%s]",
				strings.Join(*filters.Category, ","),
			),
		)
	}

	return result
}
