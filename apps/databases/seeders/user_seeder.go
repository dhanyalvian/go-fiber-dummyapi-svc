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

var userSeedData []entities.UserDoc

func init() {
	raw, err := os.ReadFile("apps/databases/data/users.json")
	if err != nil {
		log.Fatalf("Failed to read apps/databases/data/users.json: %v", err)
	}
	if err := json.Unmarshal(raw, &userSeedData); err != nil {
		log.Fatalf("Failed to parse apps/databases/data/users.json: %v", err)
	}
}

func SeedUserToTypesense(ts *typesense.Client) {
	ctx := context.Background()
	collectionName := entities.User{}.ColletionName()

	docs := make([]interface{}, len(userSeedData))
	for i, d := range userSeedData {
		docs[i] = d
	}

	databases.ImportDocuments(ts, ctx, collectionName, docs)
}
