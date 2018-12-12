package api

import (
	"net/http"

	"github.com/bli940505/threaded/server/internal/types"
	"github.com/go-chi/chi"
)

// RouteExample TODO
func RouteExample(s *types.Server) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetPerson(s))
	router.Delete("/{id}", DeletePerson(s))
	router.Post("/{id}", CreatePerson(s))
	return router
}

// GetPerson TODO
func GetPerson(s *types.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// DeletePerson TODO
func DeletePerson(s *types.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// CreatePerson TODO
func CreatePerson(s *types.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
