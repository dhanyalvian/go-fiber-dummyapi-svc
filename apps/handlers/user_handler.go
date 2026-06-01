//- apps/handlers/user_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

type UserHandler struct {
	TS *typesense.Client
}

func NewUserHandler(ts *typesense.Client) *UserHandler {
	return &UserHandler{TS: ts}
}

// @Summary List Data User
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Param query query string false "Search String"
// @Success 200 {object} response.ResponseData{data=[]entities.RespListUser}
// @Failure 400,401,404,500 {object} response.ResponseData
// @Router /user [get]
func (h *UserHandler) List(c *fiber.Ctx) error {
	docs, _ := models.ListUser(c, h.TS)
	return RespSucessList[entities.RespListUser](c, docs)
}

func (h *UserHandler) Detail(c *fiber.Ctx) error {
	id := GetId(c)
	doc, err := models.DetailUser(c, h.TS, id)
	if err != nil {
		return RespError(c, 400, "Data not found", nil)
	}

	return RespSuccessDetail[entities.RespDetailUser](c, doc)
}
