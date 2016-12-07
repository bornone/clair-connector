package cliargs

import "github.com/codegangsta/cli"

var FlAddress = cli.StringFlag{
	Name:  "listen-address",
	Usage: "<ip>:<port> to listen on",
	Value: ":8600",
}

// RegistryURL is the URI (address + endpoint) for etcd service API
var RegistryURL = cli.StringFlag{
	Name:   "registry-url",
	Usage:  "URL (http://<ip>:<port>/path) of the etcd server",
	Value:  "http://localhost:8500/v2",
	EnvVar: "REGISTRY_URL",
}

// ClairURL is the URI (address(service name in future) + endpoint) clair service API
var ClairURL = cli.StringFlag{
	Name:   "clair-url",
	Usage:  "URL (http://<ip>:<port>/path) of the clair server",
	Value:  "http://localhost:6060/v1",
	EnvVar: "CLAIR_URL",
}

var LogLevel = cli.StringFlag{
	Name:   "log-level",
	Usage:  "panic, fatal, error, warn, info, debug",
	Value:  "error",
	EnvVar: "LOG_LEVEL",
}
