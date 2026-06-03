//- apps/databases/seeders/product_seeder.go

package seeders

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"go-fiber-dummyapi-svc/apps/databases"
	"go-fiber-dummyapi-svc/apps/entities"

	"github.com/typesense/typesense-go/v4/typesense"
)

var productSeedData []entities.ProductDoc

func init() {
	raw, err := os.ReadFile("apps/databases/data/products.json")
	if err != nil {
		log.Fatalf("Failed to read apps/databases/data/products.json: %v", err)
	}
	if err := json.Unmarshal(raw, &productSeedData); err != nil {
		log.Fatalf("Failed to parse apps/databases/data/products.json: %v", err)
	}
}

func SeedProductToTypesense(ts *typesense.Client) {
	ctx := context.Background()
	collectionName := entities.Product{}.ColletionName()

	docs := make([]interface{}, len(productSeedData))
	for i, d := range productSeedData {
		docs[i] = d
	}

	databases.ImportDocuments(ts, ctx, collectionName, docs)
}