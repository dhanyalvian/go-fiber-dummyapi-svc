package handlers

import (
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

type RecipeHandler struct {
	TS *typesense.Client
}

func NewRecipeHandler(ts *typesense.Client) *RecipeHandler {
	return &RecipeHandler{TS: ts}
}

func (h *RecipeHandler) List(c *fiber.Ctx) error {
	docs, _ := models.ListRecipe(c, h.TS)
	return RespSucessList[entities.RespListRecipe](c, docs)
}

func (h *RecipeHandler) Detail(c *fiber.Ctx) error {
	id := GetId(c)
	doc, err := models.DetailRecipe(c, h.TS, id)
	if err != nil {
		return RespError(c, 400, "Data not found", nil)
	}

	return RespSuccessDetail[entities.RespDetailRecipe](c, doc)
}
