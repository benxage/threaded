package api

import (
	"net/http"

	"github.com/bli940505/threaded/server/instance"
	"github.com/go-chi/chi"
)

// RouteExample TODO
func RouteExample(in *instance.ServerInstance) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetPerson(in))
	router.Delete("/{id}", DeletePerson(in))
	router.Post("/{id}", CreatePerson(in))
	return router
}

// CreatePerson TODO
func CreatePerson(in *instance.ServerInstance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// DeletePerson TODO
func DeletePerson(in *instance.ServerInstance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// GetPerson TODO
func GetPerson(in *instance.ServerInstance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
