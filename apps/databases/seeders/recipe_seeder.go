//- apps/databases/seeders/recipe_seeder.go

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

var recipeSeedData []entities.RecipeDoc

func init() {
	raw, err := os.ReadFile("apps/databases/data/recipes.json")
	if err != nil {
		log.Fatalf("Failed to read apps/databases/data/recipes.json: %v", err)
	}
	if err := json.Unmarshal(raw, &recipeSeedData); err != nil {
		log.Fatalf("Failed to parse apps/databases/data/recipes.json: %v", err)
	}
}

func SeedRecipeToTypesense(ts *typesense.Client) {
	ctx := context.Background()
	collectionName := entities.Recipe{}.ColletionName()

	docs := make([]interface{}, len(recipeSeedData))
	for i, d := range recipeSeedData {
		docs[i] = d
	}

	databases.ImportDocuments(ts, ctx, collectionName, docs)
}
