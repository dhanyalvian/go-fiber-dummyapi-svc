//- apps/entities/product_entity.go

package entities

import (
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/shopspring/decimal"
	"github.com/typesense/typesense-go/v4/typesense/api"
)

type Product struct {
	BaseID

	Name        string `json:"name" typesense:"index,sort"`
	Description string `json:"description" typesense:"optional"`
	SKU         string `json:"sku" typesense:"index,sort"`

	Brand    string   `json:"brand" typesense:"index,facet,sort"`
	Category string   `json:"category" typesense:"index,facet,sort"`
	Tags     []string `json:"tags" typesense:"index"`

	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"images"`

	Weight    decimal.Decimal  `json:"weight"`
	Dimension ProductDimension `json:"dimensions"`

	Price    decimal.Decimal `json:"price"`
	Discount decimal.Decimal `json:"discount"`
	Stock    int             `json:"stock" typesense:"sort"`
}

type ProductDimension struct {
	Width  decimal.Decimal `json:"width"`
	Height decimal.Decimal `json:"height"`
	Depth  decimal.Decimal `json:"depth"`
}

type ProductDoc struct {
	Product
}

type RespListProduct struct {
	BaseID

	Name      string          `json:"name"`
	SKU       string          `json:"sku"`
	Brand     string          `json:"brand"`
	Category  string          `json:"category"`
	Tags      []string        `json:"tags"`
	Thumbnail string          `json:"thumbnail"`
	Price     decimal.Decimal `json:"price"`
	Discount  decimal.Decimal `json:"discount"`
	Stock     int             `json:"stock"`
}

type RespDetailProduct struct {
	RespListProduct

	Images    []string         `json:"images"`
	Weight    decimal.Decimal  `json:"weight"`
	Dimension ProductDimension `json:"dimensions"`
}

func (Product) ColletionName() string {
	return GetCollectionName(COLLECTION_PRODUCT)
}

func (Product) TypesenseSchema() ([]api.Field, *string) {
	return utils.DeriveTypesenseFields[Product](), nil
}
