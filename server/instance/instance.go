package instance

import (
	"github.com/bli940505/threaded/server/database"
	"github.com/bli940505/threaded/server/internal/config"
)

// Instance represents the backend passed around the codebase.
// Err is a reserved channel for error handling
type Instance struct {
	Config   *config.Config
	Database database.Database
	Err      chan error
}

// New TODO
func New() *Instance {
	server := &Instance{
		Err: make(chan error),
	}

	return server
}

// SetConfig TODO
func (in *Instance) SetConfig(c *config.Config) {
	in.Config = c
}

// SetDatabase TODO
func (in *Instance) SetDatabase(db database.Database) {
	in.Database = db
}
