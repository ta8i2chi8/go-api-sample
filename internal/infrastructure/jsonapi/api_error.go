package jsonapi

import (
	"fmt"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("json placeholder API error: status %d, message: %s", e.StatusCode, e.Message)
}
