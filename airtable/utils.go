package airtable

import (
	"context"
	"errors"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/mehanizm/airtable"
	"github.com/turbot/steampipe-plugin-sdk/v5/connection"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func rawConnect(ctx context.Context, connection *plugin.Connection, connectioncache *connection.ConnectionCache) (*airtable.Client, error) {
	// get airtable client from cache
	cacheKey := "airtable"
	if cachedData, ok := connectioncache.Get(ctx, cacheKey); ok {
		return cachedData.(*airtable.Client), nil
	}

	airtableConfig := GetConfig(connection)

	token := os.Getenv("AIRTABLE_TOKEN")

	if airtableConfig.Token != nil {
		token = *airtableConfig.Token
	}
	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	client := airtable.NewClient(token)

	// Save to cache
	connectioncache.Set(ctx, cacheKey, client)

	return client, nil
}

func connect(ctx context.Context, d *plugin.QueryData) (*airtable.Client, error) {
	client, err := rawConnect(ctx, d.Connection, d.ConnectionCache)
	return client, err
}

func toTableName(databaseID string, rawTableName string) string {
	return databaseID + "_" + strcase.ToSnake(rawTableName)
}
