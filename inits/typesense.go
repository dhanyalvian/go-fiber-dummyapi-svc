//- inits/typesense.go

package inits

import (
	"fmt"

	"go-fiber-dummyapi-svc/apps/configs"

	"github.com/typesense/typesense-go/v4/typesense"
)

func InitTs(cfg *configs.Config) *typesense.Client {
	cfgTs := cfg.Typesense
	server := fmt.Sprintf("%s:%d", cfgTs.Hostname, cfgTs.Port)
	apiKey := cfgTs.ApiKey

	return typesense.NewClient(
		typesense.WithServer(server),
		typesense.WithAPIKey(apiKey),
	)
}
