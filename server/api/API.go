package api

import (
	"github.com/bli940505/threaded/server/internal/types"
	"github.com/go-chi/chi"
)

// Routes returns a router with all of the Person's routes attached
func Routes(s *types.Server) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/example", RouteExample(s))

	/*************************************
	**MOUNT YOUR ROUTES HERE LIKE ABOVE!**
	*************************************/

	return router
}
