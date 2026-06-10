//- apps/entities/base_entity.go

package entities

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	SCHEMA = "dummy"

	COLLECTION_USER    = "users"
	COLLECTION_PRODUCT = "products"
	COLLECTION_RECIPE  = "recipes"
	COLLECTION_POST    = "posts"
)

type BaseID struct {
	ID string `json:"id"`
}

type BaseTimestamp struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func init() {
	decimal.MarshalJSONWithoutQuotes = true
}

func GetCollectionName(collection string) string {
	return fmt.Sprintf("%s_%s", SCHEMA, collection)
}
