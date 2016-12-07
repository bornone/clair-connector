package server

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

func about(ctx context.Context, w http.ResponseWriter, r *http.Request) *HTTPError {
	message := "Conductor for Containers Network Management API."
	message += "\n"
	message += "The listener is working. Now try a valid endpoint"
	message += "\n"
	message += "Available APIs:"
	message += "\n"
	for _, p := range endpoints {
		message += p
		message += "\n"
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%v", message)
	return nil
}
