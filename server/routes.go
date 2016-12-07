package server

import (
	"golang.org/x/net/context"
	"net/http"
)

// Handler types a function that can "handle"  http requests and responses
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) *HTTPError

// Route is a data structure to hold information of Endpoints, with its path (URI)
// and function handlers
type Route struct {
	Method          string
	Path            string
	HandlerFunction Handler
}

var endpoints []string

// Routes is an array of Routes, to store the routes (i.e., an Array of type Route)
// See Route doc for details on Route type
var Routes = []Route{
	Route{"GET", "/", about},
	Route{"GET", "/clair-connector/api/v1", about},
	Route{"POST", "/clair-connector/api/v1/event", event},
}
