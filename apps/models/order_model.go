//- apps/models/order_model.go

package models

import (
	"fmt"
	"go-fiber-dummyapi-svc/apps/dto"
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/pkgs/date"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

func ListOrder(
	c *fiber.Ctx,
	ts *typesense.Client,
	filters *dto.ReqListOrderFilter,
) (*api.SearchResult, error) {
	queryBy := "orderNo,customerName,customerEmail,status,paymentMethod"
	filterBy := getListOrdeFilters(filters)
	sortBy := []string{"orderDate:desc"}

	return GetList(c, ts, entities.Order{}.ColletionName(), queryBy, filterBy, sortBy)
}

func DetailOrder(c *fiber.Ctx, ts *typesense.Client, id string) (map[string]any, error) {
	return GetDetailById(c, ts, entities.Order{}.ColletionName(), id)
}

func getListOrdeFilters(filters *dto.ReqListOrderFilter) []string {
	var result []string

	if filters == nil {
		return result
	}

	//- filter by Order Date (DateRange)
	result = getListOrdeFilterOrderDate(filters, result)

	//- filter by Status
	result = getListOrdeFilterStatus(filters, result)

	//- filter by Payment Method
	result = getListOrdeFilterPaymentMethod(filters, result)

	return result
}

func getListOrdeFilterOrderDate(filters *dto.ReqListOrderFilter, result []string) []string {
	if filters.OrderDate != nil {
		start := &filters.OrderDate.Start
		end := &filters.OrderDate.End

		if time.Time(*end).Before(time.Time(*start)) {
			*end = *start
		}

		result = append(
			result,
			fmt.Sprintf(
				"orderDate:[`%s` TO `%s`]",
				date.FormatDateTypesense(start),
				date.FormatDateTypesense(end),
			),
		)
	}

	return result
}

func getListOrdeFilterStatus(filters *dto.ReqListOrderFilter, result []string) []string {
	if filters.Status != nil && len(*filters.Status) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"status:=[%s]",
				strings.Join(*filters.Status, ","),
			),
		)
	}

	return result
}
func getListOrdeFilterPaymentMethod(filters *dto.ReqListOrderFilter, result []string) []string {
	if filters.PaymentMethod != nil && len(*filters.PaymentMethod) > 0 {
		result = append(
			result,
			fmt.Sprintf(
				"paymentMethod:=[%s]",
				strings.Join(*filters.PaymentMethod, ","),
			),
		)
	}

	return result
}
