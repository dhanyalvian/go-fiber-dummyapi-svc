// - pkgs/request/request.go

package request

import (
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RequestLog struct {
	ReqId   string      `json:"reqId,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Query   interface{} `json:"query,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}

func GetPage(c *fiber.Ctx) int {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	return page
}

func GetLimit(c *fiber.Ctx) int {
	limit, _ := strconv.Atoi(c.Query("limit"))
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 20
	}

	return limit
}

func CountItems(v interface{}) int {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		return rv.Len()
	}

	return 0
}
