//- apps/models/order_model.go

package models

import (
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListOrder(c *fiber.Ctx, ts *typesense.Client) (*api.SearchResult, error) {
	queryBy := "orderNo,customerName,customerEmail,status,paymentMethod"
	sortBy := []string{"orderDate:desc"}
	return GetList(c, ts, entities.Order{}.ColletionName(), queryBy, "", sortBy)
}

func DetailOrder(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Order{}.ColletionName(), id)
}
