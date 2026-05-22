//- apps/handlers/user_handler.go

package handlers

import (
	"go-fiber-dummy-svc/apps/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
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
	respData := models.ListUser(c, h.DB, GetQuerySearch(c))
	return RespSucess(c, "", respData)
}

func (h *UserHandler) Detail(c *fiber.Ctx) error {
	id := c.Params("id")
	respData := models.DetailUser(h.DB, id)
	return RespSucess(c, "", respData)
}
