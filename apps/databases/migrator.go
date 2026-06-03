//- apps/databases/migrator.go

package databases

import (
	"context"
	"fmt"
	"log"

	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
	"github.com/typesense/typesense-go/v4/typesense/api/pointer"
)

type TypesenseCollection interface {
	ColletionName() string
	TypesenseSchema() ([]api.Field, *string)
}

func MigrateTypesense(ts *typesense.Client) {
	fmt.Println("Running Typesense migrations...")

	collections := []TypesenseCollection{
		entities.User{},
		entities.Product{},
		entities.Recipe{},
	}

	ctx := context.Background()

	for _, col := range collections {
		name := col.ColletionName()
		fields, defaultSort := col.TypesenseSchema()

		dropCollection(ts, ctx, name)
		createCollection(ts, ctx, name, fields, defaultSort)
	}

	fmt.Println("Typesense migrations completed successfully!")
}

func dropCollection(ts *typesense.Client, ctx context.Context, name string) {
	_, err := ts.Collection(name).Delete(ctx)
	if err != nil {
		log.Printf("[TS] Collection %s does not exist, skipping delete: %v", name, err)
		return
	}
	fmt.Printf("[TS] Collection %s dropped\n", name)
}

func createCollection(
	ts *typesense.Client,
	ctx context.Context,
	name string,
	fields []api.Field,
	defaultSort *string,
) {
	schema := &api.CollectionSchema{
		Name:                name,
		Fields:              fields,
		DefaultSortingField: defaultSort,
		EnableNestedFields:  pointer.True(),
	}

	_, err := ts.Collections().Create(ctx, schema)
	if err != nil {
		log.Fatalf("[TS] Failed to create collection %s: %v", name, err)
	}
	fmt.Printf("[TS] Collection %s created\n", name)
}
