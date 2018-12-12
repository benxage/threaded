package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bli940505/threaded/server/api"
	"github.com/bli940505/threaded/server/database"
	"github.com/bli940505/threaded/server/instance"
	"github.com/bli940505/threaded/server/internal/config"
	"github.com/bli940505/threaded/server/internal/errors"
)

func main() {
	// You can specifiy a *.toml file as the config by passing in `-env=<filename>` or `-db=<filename>`
	configFilename := flag.String("config", "config", "config filename; must exist under server/ directory")
	databaseFilename := flag.String("db", "postgres", "database filename; must exist under server/ directory")
	flag.Parse()

	// register a new server with the HandleErrors
	serverInstance := instance.New()
	errors.HandleErrors(serverInstance)

	serverInstance.SetConfig(config.New(*configFilename))
	serverInstance.SetDatabase(database.NewPostgres(*databaseFilename))

	// print config read
	serverInstance.Config.PrintInfo()
	serverInstance.Database.PrintInfo()

	// get the root router and print routes
	router := api.Router(serverInstance)
	serverInstance.Err <- api.PrintRoutes(router)

	// start server
	log.Printf("Application running on %s:%s\n", serverInstance.Config.Host, serverInstance.Config.Port)
	log.Fatal(http.ListenAndServe(":"+serverInstance.Config.Port, router))
}
