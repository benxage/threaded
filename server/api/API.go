package api

import (
	"github.com/bli940505/threaded/server/instance"
	"github.com/go-chi/chi"
)

// Routes returns a router with all of the Person's routes attached
func Routes(in *instance.Instance) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/example", RouteExample(in))

	/*************************************
	**MOUNT YOUR ROUTES HERE LIKE ABOVE!**
	*************************************/

	return router
}
