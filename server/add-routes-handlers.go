package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"github.com/bornone/clair-connector/common"
)

// AddRoutesHandlers is a function that parses the RoutesArray (array of Route)
// to create HandlerFunctions for the HTTP Server Mux
func AddRoutesHandlers(ctx context.Context, routes *[]Route) *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	for _, route := range *routes {
		localFct := route.HandlerFunction
		method := route.Method
		path := route.Path

		// wrap is the actual Handler function
		wrap := func(w http.ResponseWriter, r *http.Request) {
			//log.WithFields(log.Fields{"method": r.Method, "uri": r.RequestURI}).Info("HTTP request received")
			err := localFct(ctx, w, r)
			if err != nil {
				common.LOG(log.ErrorLevel, err.Description+" ({0} {1})", method, path)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("X-Content-Type-Options", "nosniff")
				w.WriteHeader(err.Status)
				enc := json.NewEncoder(w)
				enc.Encode(err)
				return
			}
		}

		r.Path(path).Methods(method).HandlerFunc(wrap)
		endpoints = append(endpoints, fmt.Sprintf("%v %v ", method, path))
		common.LOG(log.DebugLevel, "Registering endpoint: {0}", route)
	}
	return r
}

// Obsolete Code... remove later
// AddRoutesHandlers is a function that parses the routes map and
// adds mux Routes, for each endpoint and Handler listed in Handler map
func _AddRoutesHandlers(ctx context.Context, routes map[string]map[string]Handler) *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	for method, mappings := range routes {
		for route, fct := range mappings {
			localFct := fct

			// wrap is the actual Handler function
			wrap := func(w http.ResponseWriter, r *http.Request) {
				//log.WithFields(log.Fields{"method": r.Method, "uri": r.RequestURI}).Info("HTTP request received")
				err := localFct(ctx, w, r)
				if err != nil {
					//log.WithFields(log.Fields{"method": r.Method, "uri": r.RequestURI}).Info(err.Description)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					// w.Header().Set("X-Content-Type-Options", "nosniff")
					w.WriteHeader(err.Status)
					enc := json.NewEncoder(w)
					enc.Encode(err)
					return
				}
			}
			r.Path(route).Methods(method).HandlerFunc(wrap)
			common.LOG(log.DebugLevel, "Registering endpoint: {0}", route)
		}
	}
	return r
}
