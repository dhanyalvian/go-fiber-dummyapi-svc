//- apps/handlers/post_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

type PostHandler struct {
	TS *typesense.Client
}

func NewPostHandler(ts *typesense.Client) *PostHandler {
	return &PostHandler{TS: ts}
}

func (h *PostHandler) List(c *fiber.Ctx) error {
	docs, err := models.ListPost(c, h.TS)
	if err != nil {
		return RespError(c, 500, "Internal server error", nil)
	}

	return RespSucessList[entities.RespListPost](c, docs)
}

func (h *PostHandler) Detail(c *fiber.Ctx) error {
	id := GetId(c)
	doc, err := models.DetailPost(c, h.TS, id)
	if err != nil {
		return RespError(c, 400, "Data not found", nil)
	}

	return RespSuccessDetail[entities.RespDetailPost](c, doc)
}