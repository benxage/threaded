package types

import (
	"github.com/bli940505/threaded/server/database"
)

// Config defaults to localhost:4000
type Config struct {
	Host string
	Port string
}

// Server represents the backend passed around the codebase.
// Err is a reserved channel for error handling
type Server struct {
	Config Config
	DB     *database.Postgres
	Err    chan error
}
