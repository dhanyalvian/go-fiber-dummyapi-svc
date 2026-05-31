//- apps/models/base_model.go

package models

import (
	"go-fiber-dummyapi-svc/pkgs/request"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
	"github.com/typesense/typesense-go/v4/typesense/api/pointer"
)

func GetFilter(
	c *fiber.Ctx,
	tsClient *typesense.Client,
	tsCollection string,
	filterBy string,
) (*api.SearchResult, error) {
	querySearch := "*"
	searchParams := &api.SearchCollectionParams{
		Q:        &querySearch,
		FilterBy: pointer.String(filterBy),
		PerPage:  pointer.Int(1),
	}

	return tsClient.Collection(tsCollection).Documents().Search(c.Context(), searchParams)
}

func GetList(
	c *fiber.Ctx,
	tsClient *typesense.Client,
	tsCollection string,
	queryBy string,
) (*api.SearchResult, error) {
	querySearch := c.Query("search", "*")
	page := request.GetPage(c)
	limit := request.GetLimit(c)

	searchParams := &api.SearchCollectionParams{
		Q:       &querySearch,
		QueryBy: &queryBy,
		Page:    pointer.Int(page),
		PerPage: pointer.Int(limit),
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
