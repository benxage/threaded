package instance

import (
	"github.com/bli940505/threaded/server/database"
	"github.com/bli940505/threaded/server/internal/config"
)

// Instance represents the backend passed around the codebase.
// Err is a reserved channel for error handling
type ServerInstance struct {
	Config   *config.Config
	Database database.Database
	Err      chan error
}

// New TODO
func New() *ServerInstance {
	return &ServerInstance{Err: make(chan error)}
}

// SetConfig TODO
func (in *ServerInstance) SetConfig(c *config.Config) {
	in.Config = c
}

// SetDatabase TODO
func (in *ServerInstance) SetDatabase(db database.Database) {
	in.Database = db
}
