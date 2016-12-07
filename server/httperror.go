package server

import "net/http"

// HTTPError is
type HTTPError struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"-"`
}

// NewHTTPError is
func NewHTTPError(description string, status int) *HTTPError {
	return &HTTPError{
		Title:       http.StatusText(status),
		Description: description,
		Status:      status,
	}
}
