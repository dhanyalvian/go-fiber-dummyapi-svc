//- apps/entities/order_entity.go

package entities

import (
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/shopspring/decimal"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

type Order struct {
	BaseID

	OrderNo        string          `json:"orderNo" typesense:"index"`
	CustomerID     string          `json:"customerId" typesense:"index"`
	CustomerName   string          `json:"customerName" typesense:"index"`
	CustomerEmail  string          `json:"customerEmail" typesense:"index"`
	CustomerAvatar string          `json:"customerAvatar"`
	Status         string          `json:"status" typesense:"index,facet"`
	PaymentMethod  string          `json:"paymentMethod" typesense:"index,facet"`
	Items          []OrderItem     `json:"items"`
	TotalPrice     decimal.Decimal `json:"totalPrice"`
	TotalDiscount  decimal.Decimal `json:"totalDiscount"`
	FinalPrice     decimal.Decimal `json:"finalPrice"`
	OrderDate      string          `json:"orderDate" typesense:"sort"`
}

type OrderItem struct {
	ProductID string          `json:"productId"`
	Name      string          `json:"name"`
	SKU       string          `json:"sku"`
	Price     decimal.Decimal `json:"price"`
	Discount  decimal.Decimal `json:"discount"`
	Quantity  int             `json:"quantity"`
	Subtotal  decimal.Decimal `json:"subtotal"`
}

type OrderDoc struct {
	Order
}

type RespListOrder struct {
	BaseID

	OrderNo        string          `json:"orderNo"`
	CustomerID     string          `json:"customerId"`
	CustomerName   string          `json:"customerName"`
	CustomerEmail  string          `json:"customerEmail"`
	CustomerAvatar string          `json:"customerAvatar"`
	Status         string          `json:"status"`
	PaymentMethod  string          `json:"paymentMethod"`
	TotalPrice     decimal.Decimal `json:"totalPrice"`
	TotalDiscount  decimal.Decimal `json:"totalDiscount"`
	FinalPrice     decimal.Decimal `json:"finalPrice"`
	OrderDate      string          `json:"orderDate"`
}

type RespDetailOrder struct {
	RespListOrder

	Items []OrderItem `json:"items"`
}

func (Order) ColletionName() string {
	return GetCollectionName(COLLECTION_ORDER)
}

func (Order) TypesenseSchema() ([]api.Field, *string) {
	return utils.DeriveTypesenseFieldsWithDefaultSort[Order]()
}
