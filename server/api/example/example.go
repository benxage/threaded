package example

import (
	"fmt"
	"net/http"

	"github.com/bli940505/slackChan/server/internals/configs"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Routes returns a router with all of the Example routes attached
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetExample(configuration))
	router.Get("/", GetAllExamples(configuration))
	router.Delete("/{id}", DeleteExample(configuration))
	router.Post("/{id}", CreateExample(configuration))
	return router
}

// Example represents a returned example
type Example struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// GetExample returns a single example by id
func GetExample(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		example := Example{
			Title: "SlackChan running on PORT " + configuration.Constants.PORT,
			Body:  fmt.Sprintf("ID you gave is %s", id),
		}
		render.JSON(w, r, example) // A chi router helper for serializing and returning json
	}
}

// GetAllExamples returns all the examples stored in the database
func GetAllExamples(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		examples := []Example{
			{
				Title: "SlackChan running on PORT " + configuration.Constants.PORT,
				Body:  "I HAVE NEVER HAD SEX WITH WAYNE",
			},
			{
				Title: "SlackChan running on PORT " + configuration.Constants.PORT,
				Body:  "I HAVE NEVER HAD SEX WITH DRAKE",
			},
		}
		render.JSON(w, r, examples)
	}
}

// DeleteExample deletes an example by id from the database
func DeleteExample(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Deleted Example successfully"
		render.JSON(w, r, response)
	}
}

// CreateExample creates an example by id
func CreateExample(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Created Example successfully"
		render.JSON(w, r, response)
	}
}
