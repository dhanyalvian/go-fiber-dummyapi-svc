//- apps/models/comment_model.go

package models

import (
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListComment(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "user_id,post_id,body"
	return GetList(c, ts, entities.Comment{}.ColletionName(), queryBy)
}

func DetailComment(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Comment{}.ColletionName(), id)
}
