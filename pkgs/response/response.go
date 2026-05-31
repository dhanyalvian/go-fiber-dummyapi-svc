//- pkgs/response/response.go

package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Meta    ResponseMeta `json:"meta"`
	Message string       `json:"message"`

	Pagination *ResponsePagination `json:"pagination,omitempty"`
	Records    interface{}         `json:"records,omitempty"`
	Record     interface{}         `json:"record,omitempty"`
	Errors     interface{}         `json:"errors,omitempty"`
}

type ResponseMeta struct {
	RequestId string `json:"reqId"`
	Code      string `json:"code"`
}

type ResponsePagination struct {
	Page        int   `json:"page"`
	Next        int   `json:"next"`
	Record      int   `json:"record"`
	TotalPage   int   `json:"totalPage"`
	TotalRecord int64 `json:"totalRecord"`
}

type ResponseLog struct {
	ReqId      string      `json:"reqId,omitempty"`
	Headers    interface{} `json:"headers,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

func GetResponseReqId(c *fiber.Ctx) string {
	return string(c.Response().Header.Peek(fiber.HeaderXRequestID))
}

func GetResponseStatusCode(c *fiber.Ctx) int {
	return c.Response().StatusCode()
}
