package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bli940505/threaded/server/instance"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// PrintRoutes will walk and print all routes mounted
func PrintRoutes(router *chi.Mux) error {
	// Walk and print out all routes
	walker := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	fmt.Println("---Server Information---")
	return chi.Walk(router, walker)
}

// Router returns a router with all of the routes mounted
func Router(in *instance.ServerInstance) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Mount("/example", RouteExample(in))
	/*************************************
	**MOUNT YOUR ROUTES HERE LIKE ABOVE!**
	*************************************/

	return router
}
