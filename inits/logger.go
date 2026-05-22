//- inits/logger.go

package inits

import (
	"encoding/json"
	"log"

	"github.com/dhanyalvian/go-fiber-packages/logger"
	"github.com/dhanyalvian/go-fiber-packages/request"
	"github.com/dhanyalvian/go-fiber-packages/response"

	"github.com/gofiber/fiber/v2"
)

func InitLogger(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		initLogRequest(c)
		err := c.Next()

		initLogResponse(c)
		return err
	})
}

func initLogRequest(c *fiber.Ctx) {
	var req request.RequestLog
	req.ReqId = c.Locals("requestid").(string)
	req.Headers = c.GetReqHeaders()
	req.Params = c.AllParams()
	req.Query = c.Queries()
	req.Body = c.Body()

	logger.Logging(
		0,
		"REQUEST",
		structToJsonString(req),
	)
}

func initLogResponse(c *fiber.Ctx) {
	var resp response.ResponseLog
	var body interface{}

	resp.ReqId = c.Locals("requestid").(string)
	resp.Headers = c.GetRespHeaders()
	resp.StatusCode = c.Response().StatusCode()

	err := json.Unmarshal(c.Response().Body(), &body)
	if err == nil {
		resp.Body = body
	} else {
		resp.Body = string(c.Response().Body())
	}

	logger.Logging(
		resp.StatusCode,
		"RESPONSE",
		structToJsonString(resp),
	)
}

func structToJsonString(data any) string {
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling struct to JSON: %v", err)
	}
	return string(dataJson)
}
