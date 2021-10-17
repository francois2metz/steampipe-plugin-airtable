package airtable

import (
	"errors"
	"net/http"

	"github.com/mehanizm/airtable"
)

func isNotFoundError(err error) bool {
	responseError := &airtable.HTTPClientError{}
	if errors.As(err, &responseError) && responseError.StatusCode == http.StatusNotFound {
		return true
	}
	return false
}
