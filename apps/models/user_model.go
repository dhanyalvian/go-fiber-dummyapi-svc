//- apps/models/typesense/user_model.go

package models

import (
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListUser(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "firstname,lastname,email"
	return GetList(c, ts, entities.User{}.ColletionName(), queryBy)
}

func DetailUser(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.User{}.ColletionName(), id)
}
