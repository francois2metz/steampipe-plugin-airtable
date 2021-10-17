package airtable

import (
	"context"
	"errors"
	"net/http"

	"github.com/iancoleman/strcase"
	"github.com/mehanizm/airtable"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*airtable.Client, error) {
	// get airtable client from cache
	cacheKey := "airtable"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*airtable.Client), nil
	}

	airtableConfig := GetConfig(d.Connection)

	if &airtableConfig == nil {
		return nil, errors.New("You must have an airtable config file")
	}

	token := ""
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

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func is404Error(err error) bool {
	return isHTTPCodeError(err, http.StatusNotFound)
}

func isHTTPCodeError(err error, statusCode int) bool {
	if err == nil {
		return false
	}

	responseError := &airtable.HTTPClientError{}
	if errors.As(err, &responseError) && responseError.StatusCode == statusCode {
		return true
	}
	return false
}

func toTableName(rawTableName string) string {
	return "airtable_" + strcase.ToSnake(rawTableName)
}
