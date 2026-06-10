//- apps/handlers/comment_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

type CommentHandler struct {
	TS *typesense.Client
}

func NewCommentHandler(ts *typesense.Client) *CommentHandler {
	return &CommentHandler{TS: ts}
}

func (h *CommentHandler) List(c *fiber.Ctx) error {
	docs, err := models.ListComment(c, h.TS)
	if err != nil {
		return RespError(c, 500, "Internal server error", nil)
	}

	return RespSucessList[entities.RespListComment](c, docs)
}

func (h *CommentHandler) Detail(c *fiber.Ctx) error {
	id := GetId(c)
	doc, err := models.DetailComment(c, h.TS, id)
	if err != nil {
		return RespError(c, 400, "Data not found", nil)
	}

	return RespSuccessDetail[entities.RespDetailComment](c, doc)
}