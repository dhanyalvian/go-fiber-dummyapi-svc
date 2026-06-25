//- apps/models/quote_model.go

package models

import (
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListQuote(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "quote,author"
	sortBy := []string{"quote:asc"}
	return GetList(c, ts, entities.Quote{}.ColletionName(), queryBy, "", sortBy)
}

func DetailQuote(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Quote{}.ColletionName(), id)
}
