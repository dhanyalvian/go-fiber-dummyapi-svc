//- apps/databases/seeder.go

package databases

import (
	"context"
	"fmt"
	"log"

	"github.com/typesense/typesense-go/v4/typesense"
	"github.com/typesense/typesense-go/v4/typesense/api"
	"github.com/typesense/typesense-go/v4/typesense/api/pointer"
)

func ImportDocuments(ts *typesense.Client, ctx context.Context, name string, docs []interface{}) {
	params := &api.ImportDocumentsParams{
		Action:    pointer.Any(api.Create),
		BatchSize: pointer.Int(40),
	}

	results, err := ts.Collection(name).Documents().Import(ctx, docs, params)
	if err != nil {
		log.Fatalf("[TS] Failed to import documents: %v", err)
	}

	success := 0
	for i, r := range results {
		if r.Success {
			success++
		} else {
			fmt.Printf("[TS] Document %d failed: %s\n", i, r.Error)
		}
	}
	fmt.Printf("[TS] Imported %d/%d documents into %s\n", success, len(docs), name)
}
