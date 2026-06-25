//- apps/handlers/quote_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

type QuoteHandler struct {
	TS *typesense.Client
}

func NewQuoteHandler(ts *typesense.Client) *QuoteHandler {
	return &QuoteHandler{TS: ts}
}

func (h *QuoteHandler) List(c *fiber.Ctx) error {
	docs, err := models.ListQuote(c, h.TS)
	if err != nil {
		return RespError(c, 500, "Internal server error", nil)
	}

	return RespSucessList[entities.RespListQuote](c, docs)
}

func (h *QuoteHandler) Detail(c *fiber.Ctx) error {
	id := GetId(c)
	doc, err := models.DetailQuote(c, h.TS, id)
	if err != nil {
		return RespError(c, 400, "Data not found", nil)
	}

	return RespSuccessDetail[entities.RespDetailQuote](c, doc)
}
