//- apps/models/product_model.go

package models

import (
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListProduct(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "name,sku,brand,category,tags"
	return GetList(c, ts, entities.Product{}.ColletionName(), queryBy)
}

func DetailProduct(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Product{}.ColletionName(), id)
}
