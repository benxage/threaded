package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bli940505/slackChan/features/todo"
	"github.com/bli940505/slackChan/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Router is the main entry point that returns the project's root router
func Router(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	// TODO: mount as sub-router for version management?
	// router.Route("/v1", func(r chi.Router) {
	// 	r.Mount("/api/todo", todo.Routes(configuration))
	// })
	router.Mount("/api/todo", todo.Routes(configuration))

	return router
}

func main() {
	// You can specifiy a *.toml file as the config file by passing in `-config=<filename>`
	configFilename := flag.String("config", "config", "config filename; must exist under server folder")
	flag.Parse()

	configuration, err := config.New(*configFilename)
	if err != nil {
		log.Panicln("Configuration error", err)
	}

	router := Router(configuration)

	// Walk and print out all routes
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Println("Application running on PORT: " + configuration.Constants.PORT)
	log.Fatal(http.ListenAndServe(":"+configuration.Constants.PORT, router))
}
