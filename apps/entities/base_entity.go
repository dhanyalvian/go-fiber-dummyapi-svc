//- apps/entities/base_entity.go

package entities

import (
	"fmt"
	"time"
)

const (
	SCHEMA = "dummy"

	COLLECTION_USER = "users"
)

type BaseID struct {
	ID string `json:"id"`
}

type BaseTimestamp struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func GetCollectionName(collection string) string {
	return fmt.Sprintf("%s_%s", SCHEMA, collection)
}
