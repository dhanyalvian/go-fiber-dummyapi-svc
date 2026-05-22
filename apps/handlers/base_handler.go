// - apps/handlers/base_handler.go

package handlers

import (
	"strconv"
	"strings"

	"github.com/dhanyalvian/go-fiber-packages/response"
	"github.com/gofiber/fiber/v2"
)

func GetUid(c *fiber.Ctx) string {
	return c.Params("uid")
}

func GetUserID(c *fiber.Ctx) float64 {
	return c.Locals("user_id").(float64)
}

func GetQuerySearch(c *fiber.Ctx) string {
	result := strings.ToLower(c.Query("q"))
	return result
}

func RespSucess(c *fiber.Ctx, message string, data response.ResponseData) error {
	var resp response.Response

	resp.Meta.RequestId = response.GetResponseReqId(c)
	resp.Meta.Code = strconv.Itoa(response.GetResponseStatusCode(c))

	resp.Message = message
	resp.Data = data

	return c.Status(fiber.StatusOK).JSON(resp)
}

func RespError(c *fiber.Ctx, statusCode int, message string, err error) error {
	var resp response.Response

	resp.Meta.RequestId = response.GetResponseReqId(c)
	resp.Meta.Code = strconv.Itoa(statusCode)

	resp.Message = message
	resp.Data.Errors = err

	return c.Status(statusCode).JSON(resp)
}
