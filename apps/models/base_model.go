//- apps/models/base_model.go

package models

import (
	"github.com/dhanyalvian/go-fiber-packages/request"
	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func GetList(
	c *fiber.Ctx,
	tsClient *typesense.Client,
	tsCollection string,
	queryBy string,
) (*api.SearchResult, error) {
	querySearch := c.Query("q", "*")
	page := request.GetPage(c)
	limit := request.GetLimit(c)

	searchParams := &api.SearchCollectionParams{
		Q:       &querySearch,
		QueryBy: &queryBy,
		Page:    Ptr(page),
		PerPage: Ptr(limit),
	}

	return tsClient.Collection(tsCollection).Documents().Search(c.Context(), searchParams)
}

func GetDetailById(
	c *fiber.Ctx,
	tsClient *typesense.Client,
	tsCollection string,
	id string,
) (map[string]any, error) {
	return tsClient.Collection(tsCollection).Document(id).Retrieve(c.Context())
}

// Ptr generic helper
func Ptr[T any](v T) *T {
	return &v
}
