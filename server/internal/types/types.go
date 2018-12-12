package types

import (
	"os"

	"github.com/go-pg/pg"
)

// URL defaults to localhost:4000
type URL struct {
	Host string
	Port string
}

// Server represents the backend passed around the codebase
// Sigs is reserved channel for system signals
// Err is reserved channel for error handling
type Server struct {
	URL      URL
	Database *pg.DB
	Sigs     chan os.Signal
	Err      chan error
}
