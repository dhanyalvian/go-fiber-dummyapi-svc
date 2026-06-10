//- apps/databases/seeders/post_seeder.go

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

var postSeedData []entities.PostDoc

func init() {
	raw, err := os.ReadFile("apps/databases/data/posts.json")
	if err != nil {
		log.Fatalf("Failed to read apps/databases/data/posts.json: %v", err)
	}
	if err := json.Unmarshal(raw, &postSeedData); err != nil {
		log.Fatalf("Failed to parse apps/databases/data/posts.json: %v", err)
	}
}

func SeedPostToTypesense(ts *typesense.Client) {
	ctx := context.Background()
	collectionName := entities.Post{}.ColletionName()

	docs := make([]interface{}, len(postSeedData))
	for i, d := range postSeedData {
		docs[i] = d
	}

	databases.ImportDocuments(ts, ctx, collectionName, docs)
}