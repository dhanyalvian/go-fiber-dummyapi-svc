//- apps/databases/seeders/comment_seeder.go

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

var commentSeedData []entities.CommentDoc

func init() {
	raw, err := os.ReadFile("apps/databases/data/comments.json")
	if err != nil {
		log.Fatalf("Failed to read apps/databases/data/comments.json: %v", err)
	}
	if err := json.Unmarshal(raw, &commentSeedData); err != nil {
		log.Fatalf("Failed to parse apps/databases/data/comments.json: %v", err)
	}
}

func SeedCommentToTypesense(ts *typesense.Client) {
	ctx := context.Background()
	collectionName := entities.Comment{}.ColletionName()

	docs := make([]interface{}, len(commentSeedData))
	for i, d := range commentSeedData {
		docs[i] = d
	}

	databases.ImportDocuments(ts, ctx, collectionName, docs)
}