package seeders

import (
	"fmt"

	tsclient "github.com/typesense/typesense-go/v4/typesense"
)

func SeedAll(ts *tsclient.Client) {
	fmt.Println("Seeding Typesense...")

	SeedCommentToTypesense(ts)
	SeedPostToTypesense(ts)
	SeedProductToTypesense(ts)
	SeedRecipeToTypesense(ts)
	SeedUserToTypesense(ts)

	fmt.Println("Typesense seeding completed!")
}
