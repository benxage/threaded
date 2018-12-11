package example

import (
	"net/http"

	"github.com/bli940505/slackChan/server/internals/configs"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetSingleExample(configuration))
	router.Delete("/{id}", DeleteExample(configuration))
	router.Post("/", CreateExample(configuration))
	router.Get("/", GetAllExamples(configuration))
	return router
}

type Example struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetSingleExample(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		example := Example{
			Slug:  id,
			Title: "Hello world from PORT: " + configuration.Constants.PORT,
			Body:  "Heloo world from planet earth",
		}
		render.JSON(w, r, example) // A chi router helper for serializing and returning json
	}
}

func DeleteExample(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Deleted Example successfully"
		render.JSON(w, r, response) // Return some demo response
	}
}

func CreateExample(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Created Example successfully"
		render.JSON(w, r, response) // Return some demo response
	}
}

func GetAllExamples(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		examples := []Example{
			{
				Slug:  "slug",
				Title: "Hello world",
				Body:  "Heloo world from planet earth",
			},
		}
		render.JSON(w, r, examples) // A chi router helper for serializing and returning json
	}
}
