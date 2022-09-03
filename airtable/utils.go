package airtable

import (
	"context"
	"errors"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/mehanizm/airtable"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func rawConnect(ctx context.Context, connection *plugin.Connection) (*airtable.Client, error) {
	airtableConfig := GetConfig(connection)

	token := os.Getenv("AIRTABLE_TOKEN")
	database_id := ""

	if airtableConfig.Token != nil {
		token = *airtableConfig.Token
	}
	if airtableConfig.DatabaseID != nil {
		database_id = *airtableConfig.DatabaseID
	}
	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}
	if database_id == "" {
		return nil, errors.New("'database_id' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	client := airtable.NewClient(token)

	return client, nil
}

func connect(ctx context.Context, d *plugin.QueryData) (*airtable.Client, error) {
	// get airtable client from cache
	cacheKey := "airtable"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*airtable.Client), nil
	}

	client, err := rawConnect(ctx, d.Connection)
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func toTableName(rawTableName string) string {
	return strcase.ToSnake(rawTableName)
}
