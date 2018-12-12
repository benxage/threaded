package database

import (
	"github.com/go-pg/pg"
)

// New returns a database instance
func New(dbc *pg.Options) *pg.DB {
	return pg.Connect(dbc)
}
