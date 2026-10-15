//- apps/models/typesense/user_model.go

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

func ListUser(
	c *fiber.Ctx,
	ts *typesense.Client,
	filters *dto.ReqListUserFilter,
) (*api.SearchResult, error) {
	queryBy := "firstname,lastname,email"
	filterBy := getUserFilters(filters)
	sortBy := []string{"firstname:asc", "lastname:asc"}

	return GetList(c, ts, entities.User{}.ColletionName(), queryBy, filterBy, sortBy)
}

func getUserFilters(filters *dto.ReqListUserFilter) []string {
	var result []string

	if filters == nil {
		return result
	}

	//- filter by Gender
	result = getUserFilterGender(filters, result)

	return result
}

func getUserFilterGender(filters *dto.ReqListUserFilter, result []string) []string {
	if filters.Gender != nil && len(*filters.Gender) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"gender:=[%s]",
				strings.Join(*filters.Gender, ","),
			),
		)
	}

	return result
}

func DetailUser(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.User{}.ColletionName(), id)
}
