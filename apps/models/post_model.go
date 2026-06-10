//- apps/models/post_model.go

package models

import (
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListPost(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "user_id,title,body,tags"
	return GetList(c, ts, entities.Post{}.ColletionName(), queryBy)
}

func DetailPost(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Post{}.ColletionName(), id)
}