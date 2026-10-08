//- apps/handlers/order_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/dto"
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

type OrderHandler struct {
	TS *typesense.Client
}

func NewOrderHandler(ts *typesense.Client) *OrderHandler {
	return &OrderHandler{TS: ts}
}

func (h *OrderHandler) List(c *fiber.Ctx) error {
	filters, _ := GetQueryFilters[dto.ReqListOrderFilter](c)

	docs, err := models.ListOrder(c, h.TS, filters)
	if err != nil {
		return RespError(c, 500, "Internal server error", nil)
	}

	return RespSucessList[entities.RespListOrder](c, docs)
}

func (h *OrderHandler) Detail(c *fiber.Ctx) error {
	id := GetId(c)
	doc, err := models.DetailOrder(c, h.TS, id)
	if err != nil {
		return RespError(c, 400, "Data not found", nil)
	}

	return RespSuccessDetail[entities.RespDetailOrder](c, doc)
}
