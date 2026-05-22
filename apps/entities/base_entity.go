//- apps/entities/base_entity.go

package entities

import (
	"fmt"
	"time"
)

const (
	SCHEMA = "dummy"

	TABLE_CATEGORY = "categories"
	TABLE_PROJECT  = "projects"

	TABLE_USER = "users"
)

type BaseID struct {
	ID uint `gorm:"primaryKey;column:id" json:"id"`
}

type BaseTimestamp struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func GetTableName(table string) string {
	return fmt.Sprintf("%s.%s", SCHEMA, table)
}
