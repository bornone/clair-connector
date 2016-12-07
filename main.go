package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/bornone/clair-connector/cliargs"
	"github.com/bornone/clair-connector/common"
	"github.com/bornone/clair-connector/server"
	"github.com/codegangsta/cli"
	"golang.org/x/net/context"
)

var NAME = "clair-connector"

func main() {
	log.SetLevel(log.DebugLevel)
	common.LOG(log.DebugLevel, "Starting {0} ...", NAME)
	serveCommand := cli.Command{
		Name:      "serve",
		ShortName: "s",
		Usage:     "Serve the API",
		Flags:     []cli.Flag{cliargs.FlAddress, cliargs.RegistryURL, cliargs.ClairURL, cliargs.LogLevel},
		Action:    action(serveAction),
	}

	common.LOG(log.DebugLevel, "Command line arguments: {0}", os.Args)
	cliargs.Run(NAME, serveCommand)
}

func serveAction(c *cli.Context) error {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "listen-address", c.String("listen-address"))
	ctx = context.WithValue(ctx, "log-level", c.String("log-level"))
	ctx = context.WithValue(ctx, "clair-url", c.String("clair-url"))
	ctx = context.WithValue(ctx, "registry-url", c.String("registry-url"))

	// set log level
	ll, _ := log.ParseLevel(ctx.Value("log-level").(string))
	log.SetLevel(ll)
	return server.ServeCmd(ctx, &server.Routes)
}

func action(f func(c *cli.Context) error) func(c *cli.Context) {
	return func(c *cli.Context) {
		err := f(c)
		if err != nil {
			common.LOG(log.FatalLevel, err.Error())
		}
	}
}
