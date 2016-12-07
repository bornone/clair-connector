package server

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/bornone/clair-connector/common"
	"golang.org/x/net/context"
)

// ServeCmd does...
// TODO: comment here
func ServeCmd(ctx context.Context, routes *[]Route) error {
	common.LOG(log.DebugLevel, fmt.Sprintf("Starting HTTP server... Listening on %v...", ctx.Value("listen-address")))
	muxRouters := AddRoutesHandlers(ctx, &Routes)

	listenAddress := ctx.Value("listen-address").(string)
	return http.ListenAndServe(listenAddress, muxRouters)
}
