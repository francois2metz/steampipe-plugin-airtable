package airtable

import (
	"context"
	"errors"
	"net/http"

	"github.com/mehanizm/airtable"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func isNotFoundError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	responseError := &airtable.HTTPClientError{}
	if errors.As(err, &responseError) && responseError.StatusCode == http.StatusNotFound {
		return true
	}
	return false
}
