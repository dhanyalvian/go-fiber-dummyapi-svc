package seeders

import (
	"fmt"

	tsclient "github.com/typesense/typesense-go/v4/typesense"
)

func SeedAll(ts *tsclient.Client) {
	fmt.Println("Seeding Typesense...")

	SeedUserToTypesense(ts)
	SeedProductToTypesense(ts)
	SeedRecipeToTypesense(ts)
	SeedPostToTypesense(ts)

	fmt.Println("Typesense seeding completed!")
}
