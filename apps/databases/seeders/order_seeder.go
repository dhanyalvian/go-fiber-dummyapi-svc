//- apps/databases/seeders/order_seeder.go

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

var orderSeedData []entities.OrderDoc

func init() {
	raw, err := os.ReadFile("apps/databases/data/orders.json")
	if err != nil {
		log.Fatalf("Failed to read apps/databases/data/orders.json: %v", err)
	}
	if err := json.Unmarshal(raw, &orderSeedData); err != nil {
		log.Fatalf("Failed to parse apps/databases/data/orders.json: %v", err)
	}
}

func SeedOrderToTypesense(ts *typesense.Client) {
	ctx := context.Background()
	collectionName := entities.Order{}.ColletionName()

	docs := make([]interface{}, len(orderSeedData))
	for i, d := range orderSeedData {
		docs[i] = d
	}

	databases.ImportDocuments(ts, ctx, collectionName, docs)
}
