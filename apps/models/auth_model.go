//- apps/models/auth_model.go

package models

import (
	"fmt"
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func Login(
	c *fiber.Ctx,
	ts *typesense.Client,
	email string,
) (*api.SearchResult, error) {
	filterBy := fmt.Sprintf("email:=%s", email)
	return GetFilter(c, ts, entities.User{}.ColletionName(), filterBy)
}
