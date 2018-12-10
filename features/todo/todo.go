package todo

import (
	"net/http"

	"github.com/bli940505/slackChan/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Config TODO add comments
type Config struct {
	*config.Config
}

// New TODO add comments
func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

// Routes TODO add comments
func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", config.GetATodo)
	router.Delete("/{todoID}", config.DeleteTodo)
	router.Post("/", config.CreateTodo)
	router.Get("/", config.GetAllTodos)
	return router
}

// Todo TODO add comments
type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// GetATodo TODO add comments
func (config *Config) GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug:  todoID,
		Title: "Hello world",
		Body:  "Heloo world from planet earth",
	}
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}

// DeleteTodo TODO add comments
func (config *Config) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response) // Return some demo response
}

// CreateTodo TODO add comments
func (config *Config) CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	render.JSON(w, r, response) // Return some demo response
}

// GetAllTodos TODO add comments
func (config *Config) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}
