package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bli940505/threaded/server/api/example"
	"github.com/bli940505/threaded/server/database"
	"github.com/bli940505/threaded/server/internal/types"
	"github.com/bli940505/threaded/server/internal/utils"
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

	// TODO: mount as sub-router for version management

	router.Mount("/api/example", example.Routes(s))
	/******************************************
	**MOUNT YOUR ROUTES HERE AFTER IMPORTING!**
	******************************************/

	return router
}

func main() {
	// You can specifiy a *.toml file as the config by passing in `-env=<filename>` or `-db=<filename>`
	envFilename := flag.String("env", "env", "env filename; must exist under server/internal/configs directory")
	databaseFilename := flag.String("db", "database", "database filename; must exist under server/internal/configs directory")
	flag.Parse()

	// register a new server with the BackgroundHandler
	server := &types.Server{
		Sigs: make(chan os.Signal),
		Err:  make(chan error),
	}
	utils.BackgroundHandler(server)

	// read configs
	env, dbc, err := utils.ReadConfig(*envFilename, *databaseFilename)
	server.Err <- err
	server.URL = *env

	// utility prints for debugging
	fmt.Printf("%+v\n", *env)
	fmt.Printf("%+v\n", *dbc)

	// create new database
	db := database.New(dbc)
	server.Database = db

	// signal test
	// server.Sigs <- syscall.SIGTERM
	// server.Sigs <- syscall.SIGHUP
	// server.Sigs <- syscall.SIGINT

	// get the root router
	router := Router(server)

	// Walk and print out all routes
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	server.Err <- chi.Walk(router, walkFunc)

	log.Println(fmt.Sprintf("Application running on %s:%s\n", server.URL.Host, server.URL.Port))
	log.Fatal(http.ListenAndServe(":"+server.URL.Port, router))
}
