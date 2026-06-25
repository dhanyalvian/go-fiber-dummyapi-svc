//- apps/models/post_model.go

package models

import (
	"fmt"
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListPost(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "user_id,title,body,tags"
	sortBy := []string{"title:asc"}
	return GetList(c, ts, entities.Post{}.ColletionName(), queryBy, "", sortBy)
}

func DetailPost(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Post{}.ColletionName(), id)
}

func ListPostComment(c *fiber.Ctx, ts *typesense.Client, id string) (*api.SearchResult, error) {
	queryBy := "user_id,post_id,body"
	filterBy := fmt.Sprintf("post_id:=%s", id)
	sortBy := []string{"id:asc"}
	return GetList(c, ts, entities.Comment{}.ColletionName(), queryBy, filterBy, sortBy)
}
