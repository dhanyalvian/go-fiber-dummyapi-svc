// - apps/handlers/base_handler.go

package handlers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/dhanyalvian/go-fiber-packages/request"
	"github.com/dhanyalvian/go-fiber-packages/response"
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

func RespSucessList[T any](c *fiber.Ctx, docs *api.SearchResult) error {
	page := request.GetPage(c)
	limit := request.GetLimit(c)

	totalRecords := 0
	if docs.Found != nil {
		totalRecords = int(*docs.Found)
	}

	totalPages := 0
	if limit > 0 {
		totalPages = (totalRecords + limit - 1) / limit
	}

	next := page + 1
	if page >= totalPages {
		next = page
	}

	var results []T
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
			results = append(results, row)
		}
	}

	if results == nil {
		results = []T{}
	}

	return RespSucess(c, "", response.ResponseData{
		Pagination: &response.ResponseDataPagination{
			Page:         page,
			Next:         next,
			Records:      len(results),
			TotalPages:   totalPages,
			TotalRecords: int64(totalRecords),
		},
		Results: results,
	})
}

func RespSuccessDetail[T any](c *fiber.Ctx, doc map[string]any) error {
	var result T
	row, _ := json.Marshal(doc)
	if err := json.Unmarshal(row, &result); err != nil {
		return RespError(c, 400, "Unmarshal error", err)
	}

	return RespSucess(c, "", response.ResponseData{
		Result: result,
	})
}

func RespError(c *fiber.Ctx, statusCode int, message string, err error) error {
	var resp response.Response

	resp.Meta.RequestId = response.GetResponseReqId(c)
	resp.Meta.Code = strconv.Itoa(statusCode)

	resp.Message = message
	resp.Data.Errors = err

	return c.Status(statusCode).JSON(resp)
}
