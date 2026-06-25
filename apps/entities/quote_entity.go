//- apps/entities/quote_entity.go

package entities

import (
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/typesense/typesense-go/v4/typesense/api"
)

type Quote struct {
	BaseID

	Quote  string `json:"quote" typesense:"index,sort,defaultSort"`
	Author string `json:"author" typesense:"index,sort"`

	BaseTimestamp
}

type QuoteDoc struct {
	Quote
}

type RespListQuote struct {
	Quote
}

type RespDetailQuote struct {
	RespListQuote
}

func (Quote) ColletionName() string {
	return GetCollectionName(COLLECTION_QUOTE)
}

func (Quote) TypesenseSchema() ([]api.Field, *string) {
	return utils.DeriveTypesenseFieldsWithDefaultSort[Quote]()
}
