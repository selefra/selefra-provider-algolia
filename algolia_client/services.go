package algolia_client

import (
	"context"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func Connect(_ context.Context, config *Config) (*search.Client, error) {
	conn := search.NewClient(config.AppID, config.APIKey)
	return conn, nil
}
