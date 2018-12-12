package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bli940505/threaded/server/api"
	"github.com/bli940505/threaded/server/database"
	"github.com/bli940505/threaded/server/internal/config"
	"github.com/bli940505/threaded/server/internal/errors"
	"github.com/bli940505/threaded/server/internal/types"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Router returns the root router
func Router(s *types.Server) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Mount("/api", api.Routes(s))

	return router
}

func main() {
	// You can specifiy a *.toml file as the config by passing in `-env=<filename>` or `-db=<filename>`
	configFilename := flag.String("config", "config", "config filename; must exist under server/ directory")
	databaseFilename := flag.String("db", "database", "database filename; must exist under server/ directory")
	flag.Parse()

	// register a new server with the BackgroundHandler
	server := &types.Server{
		Err: make(chan error),
	}
	errors.HandleErrors(server)

	db, err := database.NewPostgres(*databaseFilename)
	server.Database = db
	server.Err <- err

	c, err := config.NewConfig(*configFilename)
	server.Config = c
	server.Err <- err

	// print config read
	server.Config.PrintInfo()
	server.Database.PrintInfo()

	// get the root router
	router := Router(server)

	// Walk and print out all routes
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	fmt.Println("---Server Information---")
	server.Err <- chi.Walk(router, walkFunc)

	log.Printf("Application running on %s:%s\n", server.Config.Host, server.Config.Port)
	log.Fatal(http.ListenAndServe(":"+server.Config.Port, router))
}
