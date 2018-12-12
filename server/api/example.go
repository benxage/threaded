package api

import (
	"net/http"

	"github.com/bli940505/threaded/server/instance"
	"github.com/go-chi/chi"
)

// RouteExample TODO
func RouteExample(in *instance.Instance) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetPerson(in))
	router.Delete("/{id}", DeletePerson(in))
	router.Post("/{id}", CreatePerson(in))
	return router
}

// GetPerson TODO
func GetPerson(in *instance.Instance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// DeletePerson TODO
func DeletePerson(in *instance.Instance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// CreatePerson TODO
func CreatePerson(in *instance.Instance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
