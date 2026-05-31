// - apps/handlers/base_handler.go

package handlers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"go-fiber-dummyapi-svc/pkgs/request"
	"go-fiber-dummyapi-svc/pkgs/response"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func GetId(c *fiber.Ctx) string {
	return c.Params("id")
}

func GetUid(c *fiber.Ctx) string {
	return c.Params("uid")
}

func GetUserID(c *fiber.Ctx) float64 {
	return c.Locals("user_id").(float64)
}

func GetQuerySearch(c *fiber.Ctx) string {
	result := strings.ToLower(c.Query("search"))
	return result
}

func GetDocFirst[T any](docs *api.SearchResult) T {
	var record T
	if docs.Hits != nil {
		for i, hit := range *docs.Hits {
			if hit.Document == nil {
				log.Printf("[TS] Hit[%d] document nil", i)
				continue
			}

			// var row T
			b, _ := json.Marshal(hit.Document)
			if err := json.Unmarshal(b, &record); err != nil {
				log.Printf("[TS] Hit[%d] unmarshal error: %v | raw: %s", i, err, string(b))
				continue
			}

			break
		}
	}

	return record
}

func RespSucess(
	c *fiber.Ctx,
	message string,
	record interface{},
	records interface{},
	pagination *response.ResponsePagination,
) error {
	var resp response.Response

	resp.Meta.RequestId = response.GetResponseReqId(c)
	resp.Meta.Code = strconv.Itoa(response.GetResponseStatusCode(c))

	resp.Message = message

	resp.Pagination = pagination
	resp.Records = records
	resp.Record = record

	return c.Status(fiber.StatusOK).JSON(resp)
}

func RespSucessList[T any](c *fiber.Ctx, docs *api.SearchResult) error {
	page := request.GetPage(c)
	limit := request.GetLimit(c)

	totalRecord := 0
	if docs.Found != nil {
		totalRecord = int(*docs.Found)
	}

	totalPage := 0
	if limit > 0 {
		totalPage = (totalRecord + limit - 1) / limit
	}

	next := page + 1
	if page >= totalPage {
		next = page
	}

	var records []T
	if docs.Hits != nil {
		for i, hit := range *docs.Hits {
			if hit.Document == nil {
				log.Printf("[TS] Hit[%d] document nil", i)
				continue
			}

			var row T
			b, _ := json.Marshal(hit.Document)
			if err := json.Unmarshal(b, &row); err != nil {
				log.Printf("[TS] Hit[%d] unmarshal error: %v | raw: %s", i, err, string(b))
				continue
			}
			records = append(records, row)
		}
	}

	if records == nil {
		records = []T{}
	}

	pagination := &response.ResponsePagination{
		Page:        page,
		Next:        next,
		Record:      len(records),
		TotalPage:   totalPage,
		TotalRecord: int64(totalRecord),
	}

	return RespSucess(c, "", nil, records, pagination)
}

func RespSuccessDetail[T any](c *fiber.Ctx, doc map[string]any) error {
	var record T
	row, _ := json.Marshal(doc)
	if err := json.Unmarshal(row, &record); err != nil {
		return RespError(c, 400, "Unmarshal error", err)
	}

	return RespSucess(c, "", record, nil, nil)
}

func RespError(c *fiber.Ctx, statusCode int, message string, err error) error {
	var resp response.Response

	resp.Meta.RequestId = response.GetResponseReqId(c)
	resp.Meta.Code = strconv.Itoa(statusCode)

	resp.Message = message
	resp.Errors = err

	return c.Status(statusCode).JSON(resp)
}
