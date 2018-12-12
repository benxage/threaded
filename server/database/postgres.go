package database

import (
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

// Postgres is a wrapper around pg.DB
type Postgres struct {
	Database *pg.DB
}

// NewPostgres returns a postgres instance
func NewPostgres(filename string) (*Postgres, error) {
	dbo, err := readConfig(filename)
	return &Postgres{Database: pg.Connect(dbo)}, err
}

// Insert TODO
func (db *Postgres) Insert() {

}

// Close closes the database
func (db *Postgres) Close() {
	db.Close()
}

func readConfig(filename string) (*pg.Options, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetDefault("database", "threaded")
	viper.ReadInConfig() // ignoring read error

	var dbo pg.Options
	return &dbo, viper.Unmarshal(&dbo)
}
