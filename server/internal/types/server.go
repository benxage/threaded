package types

import (
	"github.com/bli940505/threaded/server/database"
	"github.com/bli940505/threaded/server/internal/config"
)

// Server represents the backend passed around the codebase.
// Err is a reserved channel for error handling
type Server struct {
	Config   *config.Config
	Database database.Database
	Err      chan error
}
