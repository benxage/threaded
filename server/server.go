package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bli940505/threaded/server/api"
	"github.com/bli940505/threaded/server/database"
	"github.com/bli940505/threaded/server/instance"
	"github.com/bli940505/threaded/server/internal/config"
	"github.com/bli940505/threaded/server/internal/errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Router returns the root router
func Router(in *instance.Instance) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)
	router.Mount("/api", api.Routes(in))

	return router
}

func main() {
	// You can specifiy a *.toml file as the config by passing in `-env=<filename>` or `-db=<filename>`
	configFilename := flag.String("config", "config", "config filename; must exist under server/ directory")
	databaseFilename := flag.String("db", "database", "database filename; must exist under server/ directory")
	flag.Parse()

	// register a new server with the BackgroundHandler
	serverInstance := instance.New()
	errors.HandleErrors(serverInstance)

	serverInstance.SetConfig(config.New(*configFilename))
	serverInstance.SetDatabase(database.NewPostgres(*databaseFilename))

	// print config read
	serverInstance.Config.PrintInfo()
	serverInstance.Database.PrintInfo()

	// get the root router
	router := Router(serverInstance)

	// Walk and print out all routes
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	fmt.Println("---Server Information---")
	serverInstance.Err <- chi.Walk(router, walkFunc)
	log.Printf("Application running on %s:%s\n", serverInstance.Config.Host, serverInstance.Config.Port)
	log.Fatal(http.ListenAndServe(":"+serverInstance.Config.Port, router))
}
