//- apps/handlers/product_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

type ProductHandler struct {
	TS *typesense.Client
}

func NewProductHandler(ts *typesense.Client) *ProductHandler {
	return &ProductHandler{TS: ts}
}

func (h *ProductHandler) List(c *fiber.Ctx) error {
	docs, err := models.ListProduct(c, h.TS)
	if err != nil {
		return RespError(c, 500, "Internal server error", nil)
	}

	return RespSucessList[entities.RespListProduct](c, docs)
}

func (h *ProductHandler) Detail(c *fiber.Ctx) error {
	id := GetId(c)
	doc, err := models.DetailProduct(c, h.TS, id)
	if err != nil {
		return RespError(c, 400, "Data not found", nil)
	}

	return RespSuccessDetail[entities.RespDetailProduct](c, doc)
}
