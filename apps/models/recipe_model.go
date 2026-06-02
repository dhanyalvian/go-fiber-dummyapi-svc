package models

import (
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListRecipe(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "name,cuisine,tags,mealType"
	return GetList(c, ts, entities.Recipe{}.ColletionName(), queryBy)
}

func DetailRecipe(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Recipe{}.ColletionName(), id)
}
