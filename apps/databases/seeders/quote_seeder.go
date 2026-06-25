//- apps/databases/seeders/quote_seeder.go

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

var quoteSeedData []entities.QuoteDoc

func init() {
	raw, err := os.ReadFile("apps/databases/data/quotes.json")
	if err != nil {
		log.Fatalf("Failed to read apps/databases/data/quotes.json: %v", err)
	}
	if err := json.Unmarshal(raw, &quoteSeedData); err != nil {
		log.Fatalf("Failed to parse apps/databases/data/quotes.json: %v", err)
	}
}

func SeedQuoteToTypesense(ts *typesense.Client) {
	ctx := context.Background()
	collectionName := entities.Quote{}.ColletionName()

	docs := make([]interface{}, len(quoteSeedData))
	for i, d := range quoteSeedData {
		docs[i] = d
	}

	databases.ImportDocuments(ts, ctx, collectionName, docs)
}
