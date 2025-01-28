package main

import (
	"github.com/sev-2/raiden"
	"medpoint/internal/bootstrap"
)

func main() {
	// load configuration
	config, err := raiden.LoadConfig(nil)
	if err != nil {
		raiden.Error("load configuration",err.Error())
	}

	// Setup server
	server := raiden.NewServer(config)

	// register resource
	bootstrap.RegisterRoute(server)
	bootstrap.RegisterJobs(server)
	bootstrap.RegisterSubscribers(server)

	// run server
	server.Run()
}
